package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/joho/godotenv"

	"lamode.com/collector"
)

func prompt(version string, addr string) {
	promptString := `
       ┓┓        
    ┏┏┓┃┃┏┓┏╋┏┓┏┓
    ┗┗┛┗┗┗ ┗┗┗┛┛  v%s
    - Serving at http://%s
    - To close connection CTRL+C
    `
	fmt.Printf(promptString, version, addr)
	fmt.Println()
}

func run(
	ctx context.Context,
	getenv func(string) string,
	stdout io.Writer,
	stderr io.Writer,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	config := &collector.Config{
		Version: getenv("COLLECTOR_VERSION"),
		Host:    getenv("COLLECTOR_HOST"),
		Port:    getenv("COLLECTOR_PORT"),
	}

	producerLogger := log.New(stdout, "PRODUCER: ", log.Ltime)
	producer, err := collector.NewProducer(
		producerLogger,
		getenv("KAFKA_HOST"),
		getenv("KAFKA_PORT"),
		getenv("KAFKA_TOPIC"),
	)

	if err != nil {
		fmt.Fprintf(stderr, "error occured while creating producer: %s\n", err)
	}

	serverLogger := log.New(stdout, "SERVER: ", log.Ltime)
	srv := collector.NewServer(serverLogger, config, producer)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}

	go func() {
		prompt(config.Version, httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(stderr, "error serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(stderr, "error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()

	return nil
}

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error occured: %s\n", err)
		os.Exit(1)
	}

	if err := run(ctx, os.Getenv, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error occured: %s\n", err)
		os.Exit(1)
	}
}
