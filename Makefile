USER = $$(id -u)
PORT = 3001
DOCKER_COMPOSE ?= USER=$(USER) PORT=$(PORT) docker compose

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help fmt run
.DEFAULT_GOAL := help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

up: ## Run server
	$(DOCKER_COMPOSE) --profile serve up -d --remove-orphans	
	
stop: ## Stop server
	$(DOCKER_COMPOSE) --profile serve stop

restart: ## Restart server
	$(DOCKER_COMPOSE) --profile serve restart

down: stop
	$(DOCKER_COMPOSE) down --remove-orphans

build:
	$(DOCKER_COMPOSE) build grpc 
	
remove: down _image_remove _container_remove	

_image_remove:
	docker image rm -f \
		auth-grpc 

_container_remove:
	docker rm -f \
        auth_grpc