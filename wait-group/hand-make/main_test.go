package main

import (
	"testing"
	"time"
)

func TestNewWaitGroupHandMake(t *testing.T) {
	wg := NewWaitGroupHandMake()
	if wg == nil {
		t.Fatal("NewWaitGroupHandMake returned nil")
	}
	if wg.count != 0 {
		t.Errorf("Expected count to be 0, got %d", wg.count)
	}
}

func TestAdd(t *testing.T) {
	wg := NewWaitGroupHandMake()

	// Test adding positive value
	wg.Add(3)
	if wg.count != 3 {
		t.Errorf("Expected count to be 3, got %d", wg.count)
	}

	// Test adding zero
	wg.Add(0)
	if wg.count != 3 {
		t.Errorf("Expected count to remain 3, got %d", wg.count)
	}

	// Test adding negative value (should panic)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when adding negative value")
		}
	}()
	wg.Add(-4) // This should cause count to become -1 and panic
}

func TestDone(t *testing.T) {
	wg := NewWaitGroupHandMake()
	wg.Add(2)

	wg.Done()
	if wg.count != 1 {
		t.Errorf("Expected count to be 1 after Done(), got %d", wg.count)
	}

	wg.Done()
	if wg.count != 0 {
		t.Errorf("Expected count to be 0 after second Done(), got %d", wg.count)
	}
}

func TestWait(t *testing.T) {
	wg := NewWaitGroupHandMake()

	// Test Wait with zero count (should return immediately)
	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)
	if elapsed > 100*time.Millisecond {
		t.Errorf("Wait() took too long with zero count: %v", elapsed)
	}

	// Test Wait with non-zero count
	wg.Add(1)
	done := make(chan bool)

	go func() {
		wg.Wait()
		done <- true
	}()

	// Wait should block
	select {
	case <-done:
		t.Error("Wait() returned before Done() was called")
	case <-time.After(100 * time.Millisecond):
		// Expected behavior
	}

	// Call Done to unblock
	wg.Done()

	// Wait should now complete
	select {
	case <-done:
		// Expected behavior
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait() did not return after Done() was called")
	}
}

func TestMultipleGoroutines(t *testing.T) {
	wg := NewWaitGroupHandMake()
	numGoroutines := 5
	wg.Add(numGoroutines)

	completed := make(chan bool, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond) // Simulate work
			completed <- true
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Verify all goroutines completed
	close(completed)
	count := 0
	for range completed {
		count++
	}

	if count != numGoroutines {
		t.Errorf("Expected %d goroutines to complete, got %d", numGoroutines, count)
	}
}

func TestConcurrentAddAndDone(t *testing.T) {
	wg := NewWaitGroupHandMake()
	numOperations := 100

	// Start with some count
	wg.Add(numOperations / 2)

	// Concurrently add and done
	for i := 0; i < numOperations/2; i++ {
		go func() {
			wg.Add(1)
			time.Sleep(1 * time.Millisecond)
			wg.Done()
		}()
	}

	// Complete the initial count
	for i := 0; i < numOperations/2; i++ {
		wg.Done()
	}

	// This should complete without deadlock
	wg.Wait()
}

func TestBroadcastOnZero(t *testing.T) {
	wg := NewWaitGroupHandMake()
	wg.Add(2)

	waiters := make([]chan bool, 3)
	for i := range waiters {
		waiters[i] = make(chan bool)
		go func(ch chan bool) {
			wg.Wait()
			ch <- true
		}(waiters[i])
	}

	// Give waiters time to start
	time.Sleep(10 * time.Millisecond)

	// All waiters should be blocked
	for i, ch := range waiters {
		select {
		case <-ch:
			t.Errorf("Waiter %d returned before count reached zero", i)
		default:
			// Expected
		}
	}

	// Complete all work
	wg.Done()
	wg.Done()

	// All waiters should now complete
	for i, ch := range waiters {
		select {
		case <-ch:
			// Expected
		case <-time.After(100 * time.Millisecond):
			t.Errorf("Waiter %d did not return after count reached zero", i)
		}
	}
}
