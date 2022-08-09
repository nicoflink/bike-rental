package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/nicoflink/bike-rental/pkg/http/rest"
	"github.com/nicoflink/bike-rental/pkg/list"
	"github.com/nicoflink/bike-rental/pkg/persistence/memory"
)

func main() {
	log.Println("starting bike rental server")

	// Init Repos
	memoryRepo := memory.NewRepository(memory.WithSampleBikes(memory.SampleBikes))

	// Init Services
	listService := list.NewService(memoryRepo)

	server := rest.NewServer(
		"localhost",
		8080,
		nil,
		rest.DomainServices{ListService: listService},
	)

	serverCtx, serverStop := context.WithCancel(context.Background())

	// CHI EXAMPLE GRACEFUL SHUTDOWN START

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStop()
	}()

	// CHI EXAMPLE END

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.Serve(serverCtx)
		if err != nil {
			log.Printf("ERROR: %s", err)
		}
	}()

	wg.Wait()
}
