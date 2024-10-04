# Collector

## Usage
```sh
$ go run . -help
Usage of collector:
  -addr string
        The address to bind to (default ":50051")
  -env string
        Environment of program (default "dev")
  -kafka-addr string
        Kafka bootstrap servers (default ":9092")
  -kafka-topic string
        Kafka topic (default "clickstream-test")
  -verbose
        Turn on some debugging logs

# You can also use env variables
$ APP_ADDR=":8080" ENV="prod" KAFKA_ADDR=":9092" go run .
```

## Build
```sh
docker build -t <CONTAINER_TAG> ./
```

## Test
```sh
$ go test -v .
```
