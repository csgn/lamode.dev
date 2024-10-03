# Collector

## Build
```sh
$ go mod download
$ go mod verify
$ go build -v -o ./build/collector ./cmd/collector
```

## Run
```sh
$ COLLECTOR_VERSION=$(cat version) \
  COLLECTOR_HOST=<HOST> \
  COLLECTOR_PORT=<PORT> \
  KAFKA_HOST=<KAFKA_HOST> \
  KAFKA_PORT=<KAFKA_PORT> \
  KAFKA_TOPIC=<KAFKA_TOPIC> \
  ./build/collector
```

## Test
```sh
$ go test -v .
```
