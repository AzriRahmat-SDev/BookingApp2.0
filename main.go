package main

import (
	"GoInActionAssignment/database"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	port            = 5221
	idleTimeout     = 5 * time.Minute
	shutdownTimeout = 10 * time.Second
)

func init() {
	var wg sync.WaitGroup
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Lshortfile | log.Ltime | log.Lshortfile)
	wg.Add(2)
	go func() {
		database.InitDoctors()
		wg.Done()
	}()
	go func() {
		database.InitCustomers()
		wg.Done()
	}()
	wg.Wait()
}
func main() {
	server := &http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		Handler:     routes(),
		IdleTimeout: idleTimeout,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	//Communication security using HTTPS with certs and keys
	go func() {
		log.Printf("Listening on port:%d", port)
		if err := server.ListenAndServeTLS(
			"/Users/azri-rahmat/Desktop/GoSchool/Assignment/BookingApp2.0/cert.pem",
			"/Users/azri-rahmat/Desktop/GoSchool/Assignment/BookingApp2.0/key.pem"); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	// blocks code, waits for stop to initiate
	<-stop

	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server shut down")
}
