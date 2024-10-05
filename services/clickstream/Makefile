NAME := clickstream
TAG=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
PROJECT_NAME := $(NAME)-$(TAG)
PROJECT_DIR := $(PWD)
DEPLOYMENT_DIR := $(PROJECT_DIR)/docker
MODES_DIR := $(DEPLOYMENT_DIR)/modes
CMD := docker compose
COMPOSE_ENTRYPOINT := $(DEPLOYMENT_DIR)/compose.yaml
PRE_FLAGS := --file $(COMPOSE_ENTRYPOINT) 				\
			 --project-name $(PROJECT_NAME) 			\
			 --project-directory $(DEPLOYMENT_DIR)		\
			 --env-file $(MODES_DIR)/$(MODE)
POST_FLAGS := --build --detach

all:
ifndef MODE
	$(error MODE not defined. Example Usage: MODE=dev make up)
endif
	@echo "Using environment variables:"
	@echo $(shell cat $(MODES_DIR)/$(MODE))

up: all
	$(CMD) $(PRE_FLAGS) up $(POST_FLAGS)

down: all
	$(CMD) $(PRE_FLAGS) down

stop: all
	$(CMD) $(PRE_FLAGS) stop

config: all
	$(CMD) $(PRE_FLAGS) config
