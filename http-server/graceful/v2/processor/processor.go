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

func NewHttpServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
}

func publisher(ctx context.Context, conn *nats.Conn) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1 * time.Second):
			log.Println("[HTTP] Sending message")
			conn.Publish("my.subject", []byte("Hello from publisher"))
		}
	}
}

func main() {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	defer conn.Close()

	processor := NewProcessor(conn)
	httpServer := NewHttpServer()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		log.Println("[HTTP] Server is running")
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Failed to start HTTP server:", err)
		}
		log.Println("[HTTP] Server stopped")
	}()

	if err := processor.Start(5, "my.subject"); err != nil {
		log.Fatal("Failed to start processor:", err)
	}
	log.Println("[NATS] Processor started")

	go publisher(ctx, conn)

	log.Println("[HTTP] Press Ctrl+C to stop")

	<-ctx.Done()

	log.Println("[APP] Stopping...")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		log.Println("[HTTP] Stopping HTTP server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal("Failed to shutdown HTTP server:", err)
		}

		log.Println("[HTTP] HTTP server stopped")
	}()

	go func() {
		defer wg.Done()
		log.Println("[NATS] Stopping processor...")
		processor.Stop()
		log.Println("[NATS] Processor stopped")
	}()

	log.Println("[APP] Stopped")
}
