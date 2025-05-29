package day13

import (
	"context"
	"fmt"
	"go-practice/day13/trace"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Day13 struct{}

func service(ctx context.Context) {
	tid := trace.GetTraceID(ctx)
	log.Printf("[TRACE-ID %s] handling service\n", tid)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tid := trace.GetTraceID(ctx)
	log.Printf("[TRACE-ID %s] handling request\n", tid)
	service(r.Context())
	fmt.Fprintf(w, "ok\n")
}

func (d *Day13) traceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tid := r.Header.Get("X-Trace-ID")
		if tid == "" {
			tid = uuid.NewString()
		}

		ctx := trace.WithTraceID(r.Context(), tid)
		w.Header().Set("X-Trace-ID", tid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (d *Day13) Exec() {
	http.Handle("/", d.traceMiddleware(http.HandlerFunc(handler)))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
