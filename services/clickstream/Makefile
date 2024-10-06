NAME := clickstream
TAG=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
PROJECT_NAME := $(NAME)-$(TAG)
PROJECT_DIR := $(PWD)
DEPLOYMENT_DIR := $(PROJECT_DIR)/docker
MODES_DIR := $(DEPLOYMENT_DIR)/modes
CMD := docker compose
COMPOSE_ENTRYPOINT := $(DEPLOYMENT_DIR)/compose.yaml

modes := $(shell find $(MODES_DIR) -maxdepth 1 -type f -iname "*.env" -exec basename {} .env \; | sed ':a;N;$$!ba;s/\n/ | /g')

ifndef MODE
    $(error MODE does not specified. Example Usage: MODE=($(modes)) make up)
else
    MODE_FILE := $(MODES_DIR)/$(MODE).env
endif

ifeq ("$(wildcard $(MODE_FILE))", "")
    $(error '$(MODE)' does not exist. Try ($(modes)))
endif

PRE_FLAGS := --file $(COMPOSE_ENTRYPOINT) 				\
			 --project-name $(PROJECT_NAME) 			\
			 --project-directory $(DEPLOYMENT_DIR)		\
			 --env-file $(MODE_FILE)
POST_FLAGS := --build --detach


up:
	$(CMD) $(PRE_FLAGS) up $(POST_FLAGS)

down:
	$(CMD) $(PRE_FLAGS) down

stop:
	$(CMD) $(PRE_FLAGS) stop

config:
	$(CMD) $(PRE_FLAGS) config
