package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		fmt.Println("Server is running")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Errorf("error when listening server"))
		}

		fmt.Println("Server stopped")
	}()

	go func(ctx context.Context) {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Interrupt signal received")
				return
			case t := <-ticker.C:
				fmt.Println("Working:", t)
			}
		}
	}(ctx)

	fmt.Println("App running")
	<-ctx.Done()

	fmt.Println("App shutting down")

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(fmt.Errorf("error when shut down the server %v", err))
	}

	fmt.Println("Shut down completed")
}
