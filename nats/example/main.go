package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	// In docker-compose, the nats service will be available at "nats:4222"
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	var wg sync.WaitGroup
	wg.Add(3) // Wait for 3 subscribers to finish (conceptually, though they run forever here until we stop)

	// Start 3 Subscribers
	for i := 1; i <= 3; i++ {
		id := i
		go func() {
			// Subscribe to "updates"
			_, err := nc.Subscribe("updates", func(m *nats.Msg) {
				fmt.Printf("[Subscriber %d] Received: %s\n", id, string(m.Data))

				// Reply if Reply subject is present
				if m.Reply != "" {
					replyMsg := fmt.Sprintf("Reply from Subscriber %d", id)
					nc.Publish(m.Reply, []byte(replyMsg))
					fmt.Printf("[Subscriber %d] Sent reply: %s\n", id, replyMsg)
				}
			})
			if err != nil {
				log.Printf("Subscriber %d error: %v", id, err)
			}
			// Signal ready (in a real app we might want better coordination)
			// For this example, we just let them run.
		}()
	}

	// Give subscribers a moment to be ready
	time.Sleep(1 * time.Second)

	// Start Publisher
	go func() {
		// Create a unique inbox for replies
		replySubject := nats.NewInbox()

		// Subscribe to the reply subject to get responses
		sub, err := nc.SubscribeSync(replySubject)
		if err != nil {
			log.Fatal(err)
		}

		msg := "Hello NATS!"
		fmt.Printf("[Publisher] Publishing: %s\n", msg)

		// Publish with reply subject
		err = nc.PublishRequest("updates", replySubject, []byte(msg))
		if err != nil {
			log.Fatal(err)
		}

		// Wait for 3 replies
		for i := 0; i < 3; i++ {
			msg, err := sub.NextMsg(2 * time.Second)
			if err != nil {
				log.Printf("[Publisher] Error waiting for reply: %v", err)
				break
			}
			fmt.Printf("[Publisher] Received reply: %s\n", string(msg.Data))
		}
		wg.Done() // Signal we are done for the example purpose
	}()

	// Keep the app running for a bit to see output, or wait for a signal
	// For this example, we'll wait for the publisher flow to complete + a bit buffer
	// But since wg.Done is only called by publisher (incorrect usage above for full app lifecycle),
	// let's just wait a fixed time or use a channel.

	// Better approach: Wait for publisher to finish its job.
	// We reused wg for convenience but let's fix the logic.
	// We want to exit after the demo flow is done.

	// Let's just sleep enough for the demo
	time.Sleep(5 * time.Second)
	fmt.Println("Demo finished.")
}
