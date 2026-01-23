package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type CBFn func() error

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	state State

	failureCount int
	successCount int

	failureThreshold         int
	halfOpenSuccessThreshold int

	openTimeout     time.Duration
	lastFailureTime time.Time

	mu sync.Mutex
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()

	switch cb.state {
	case Open:
		if time.Since(cb.lastFailureTime) < cb.openTimeout {
			cb.mu.Unlock()
			return errors.New("circuit breaker is open")
		}
		// chuyển sang half-open
		cb.state = HalfOpen
		cb.successCount = 0
	}

	cb.mu.Unlock()

	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.onFailure()
		return err
	}

	cb.onSuccess()
	return nil
}

func (cb *CircuitBreaker) onFailure() {
	cb.failureCount++
	cb.lastFailureTime = time.Now()

	switch cb.state {
	case Closed:
		if cb.failureCount >= cb.failureThreshold {
			cb.state = Open
		}
	case HalfOpen:
		cb.state = Open
	}
}

func (cb *CircuitBreaker) onSuccess() {
	switch cb.state {
	case Closed:
		cb.failureCount = 0
	case HalfOpen:
		cb.successCount++
		if cb.successCount >= cb.halfOpenSuccessThreshold {
			cb.state = Closed
			cb.failureCount = 0
		}
	}
}

func callPaymentService() error {
	fmt.Println("payment service called")

	return nil
}

func main() {

	cb := &CircuitBreaker{
		state:                    Closed,
		failureThreshold:         5,
		halfOpenSuccessThreshold: 3,
		openTimeout:              10 * time.Second,
	}

	if err := cb.Execute(func() error {
		return callPaymentService()
	}); err != nil {
		panic(err)
	}

	fmt.Println("DONE")
}
