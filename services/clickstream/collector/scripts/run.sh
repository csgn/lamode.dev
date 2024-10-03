#!/bin/sh

COLLECTOR_VERSION=$(cat version) \
COLLECTOR_HOST=0.0.0.0 \
COLLECTOR_PORT=9999 \
KAFKA_HOST=0.0.0.0 \
KAFKA_PORT=9092 \
KAFKA_TOPIC=clickstream-test \
go run ./cmd/collector
