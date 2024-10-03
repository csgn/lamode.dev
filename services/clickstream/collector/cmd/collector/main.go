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

	"lamode.com/collector"
)

func getenv(key string, shouldExist bool) string {
	value := os.Getenv(key)

	if value == "" && shouldExist {
		panic(key + " is not defined")
	}

	return value
}

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
	getenv func(string, bool) string,
	stdout io.Writer,
	stderr io.Writer,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	config := &collector.Config{
		Version: getenv("COLLECTOR_VERSION", true),
		Host:    getenv("COLLECTOR_HOST", true),
		Port:    getenv("COLLECTOR_PORT", true),
	}

	producerLogger := log.New(stdout, "PRODUCER: ", log.Ltime)
	producer, err := collector.NewProducer(
		producerLogger,
		getenv("KAFKA_HOST", true),
		getenv("KAFKA_PORT", true),
		getenv("KAFKA_TOPIC", true),
	)
	defer producer.Close()

	if err != nil {
		return err
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

	if err := run(ctx, getenv, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error occured: %s\n", err)
		os.Exit(1)
	}
}
