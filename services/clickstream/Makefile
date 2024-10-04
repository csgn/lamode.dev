TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

HOST_NAME := clickstream
PROJECT_NAME := $(HOST_NAME)-$(TAG)
DEPLOYMENT_DIR := $(PWD)/docker
CMD := docker compose
COMPOSE_ENTRYPOINT := $(DEPLOYMENT_DIR)/compose.yaml
PRE_FLAGS := -f $(COMPOSE_ENTRYPOINT) -p $(PROJECT_NAME) --project-directory $(DEPLOYMENT_DIR) 
POST_FLAGS := -d


export KAFKA_HOST := kafka
export KAFKA_PORT := 9092
export KAFKA_ADDR := $(KAFKA_HOST):$(KAFKA_PORT)
export KAFKA_TOPIC := "$(PROJECT_NAME)-$(ENV)-topic"

export ENV ?= dev

ifeq ($(ENV), dev)
	export VERBOSE := 1
else
	export VERBOSE := 0
endif

export APP_HOST := localhost
export APP_PORT := 50051
export APP_ADDR := $(APP_HOST):$(APP_PORT)

up:
	@echo "Exported Variables:"
	@echo "KAFKA_HOST = $(KAFKA_HOST)"
	@echo "KAFKA_PORT = $(KAFKA_PORT)"
	@echo "KAFKA_ADDR = $(KAFKA_ADDR)"
	@echo "KAFKA_TOPIC = $(KAFKA_TOPIC)"
	@echo "ENV = $(ENV)"
	@echo "APP_ADDR = $(APP_ADDR)"
	@echo "APP_PORT = $(APP_PORT)"
	@echo "VERBOSE = $(VERBOSE)"
	$(CMD) $(PRE_FLAGS) up $(POST_FLAGS)

down:
	$(CMD) $(PRE_FLAGS) down

stop:
	$(CMD) $(PRE_FLAGS) stop
