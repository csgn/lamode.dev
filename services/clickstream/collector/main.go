package main

import (
	"flag"
	"log"
	"os"
	"strconv"
)

const (
	defaultEnv        = "dev"
	defaultVerbose    = false
	defaultAppAddr    = ":8080"
	defaultKafkaAddr  = ":9092"
	defaultKafkaTopic = "clickstream-test"
)

var (
	addr = flag.String(
		"addr",
		getenv("APP_ADDR", defaultAppAddr),
		"The address to bind to",
	)
	kafkaAddr = flag.String(
		"kafka-addr",
		getenv("KAFKA_ADDR", defaultKafkaAddr),
		"Kafka bootstrap servers",
	)
	kafkaTopic = flag.String(
		"kafka-topic",
		getenv("KAFKA_TOPIC", defaultKafkaTopic),
		"Kafka topic",
	)
	env = flag.String(
		"env",
		getenv("ENV", defaultEnv),
		"Environment of Collector",
	)
	verbose = flag.Bool(
		"verbose",
		getenv("VERBOSE", defaultVerbose),
		"Turn on some debugging logs",
	)
)

func getenv[T any](key string, defaultValue T) T {
	var result any

	if value, ok := os.LookupEnv(key); ok {
		switch any(defaultValue).(type) {
		case string:
			result = value
		case int:
			parsed, err := strconv.Atoi(value)
			if err != nil {
				log.Fatalf("Parse error: %s\n", key)
			}
			result = parsed
		case bool:
			parsed, err := strconv.ParseBool(value)
			if err != nil {
				log.Fatalf("Parse error: %s\n", key)
			}

			result = parsed
		default:
			return defaultValue
		}

		return result.(T)
	}

	return defaultValue
}

func main() {
	flag.Parse()

	if *addr == "" || *kafkaAddr == "" || *kafkaTopic == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	s := &Server{
		Addr:          *addr,
		AsyncProducer: NewProducer(*kafkaAddr),
	}

	defer func() {
		if err := s.Close(); err != nil {
			log.Printf("error shutting down http server: %s\n", err)
		}
	}()

	log.Fatal(s.Run())
}
