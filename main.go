package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handlers "github.com/AFORANURAG/microservices-golang/handlers"
)

func main() {
	serveMux := http.NewServeMux()
	server := http.Server{Addr: "localhost:9000", Handler: serveMux, IdleTimeout: 120 * time.Second, ReadTimeout: 60 * time.Second, WriteTimeout: 60 * time.Second}

	helloHandler := handlers.NewHelloHandler()
	goodbyeHandler := handlers.NewGoodByeHandler()
	productHandler := handlers.NewProductHandler()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)
	serveMux.Handle("/product", productHandler)
	go func() {
		// blocking operation
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
		}

	}()
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, []os.Signal{os.Interrupt, os.Kill}...)
	sig := <-sigChan
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	log.Println("Received shutdown signal:", sig)
	log.Println("Shutting down gracefully...")
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}
	log.Println("Server stopped cleanly")
	defer cancel()
}
