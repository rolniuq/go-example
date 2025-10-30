package processor

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

type Processor struct {
	conn     *nats.Conn
	sub      *nats.Subscription
	workChan chan *nats.Msg
	wg       sync.WaitGroup
}

func NewProcessor(conn *nats.Conn) *Processor {
	return &Processor{
		conn:     conn,
		workChan: make(chan *nats.Msg, 100),
	}
}

func (p *Processor) worker(id int) {
	log.Printf("[NATS] Worker %d started", id)
	for msg := range p.workChan {
		p.wg.Add(1)

		log.Println("[NATS] Worker", id, "processing message", string(msg.Data))
		time.Sleep(2 * time.Second)
		log.Println("[NATS] Worker", id, "processed message completed")

		p.wg.Done()
	}

	log.Printf("[NATS] Worker %d stopped", id)
}

func (p *Processor) natsReceiver(msg *nats.Msg) {
	p.workChan <- msg
}

func (p *Processor) Start(workerCount int, subject string) error {
	for i := range workerCount {
		go p.worker(i)
	}

	var err error
	p.sub, err = p.conn.Subscribe(subject, p.natsReceiver)
	return err
}

func (p *Processor) Stop() {
	log.Println("[NATS] Stopping processor...")
	if p.sub == nil {
		log.Println("[NATS] Processor already stopped")
		return
	}

	if err := p.sub.Drain(); err != nil {
		log.Fatalf("[NATS] Failed to drain subscription: %v", err)
		return
	}

	close(p.workChan)
	log.Println("[NATS] Processor stopped")

	p.wg.Wait()
	log.Println("[NATS] All workers stopped")
}
