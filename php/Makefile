PHP_VERSION ?= 8.4
GENERATE_DIR ?= src/generate/php
USER = $$(id -u)

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
.DEFAULT_GOAL := help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

composer: ## composer install
	docker run --init -it --rm -u ${USER} -v "$$(pwd):/app" \
		-v ".auth.json:/tmp/auth.json" \
		-w /app \
		composer:latest \
		composer install --optimize-autoloader --ignore-platform-reqs

composer-up: ## composer update
	docker run --init -it --rm -u ${USER} -v "$$(pwd):/app" \
		-v ".auth.json:/tmp/auth.json" \
		-w /app \
		composer:latest \
		composer update --no-cache --ignore-platform-reqs

composer-dump: ## composer dump-autoload
	docker run --init -it --rm -u ${USER} -v "$$(pwd):/app" -w /app \
		composer:latest \
		composer dump-autoload

composer-cli: ## composer console
	docker run --init -it --rm -u ${USER} -v "$$(pwd):/app" \
		-v ".auth.json:/tmp/auth.json" \
		-w /app \
		composer:latest \
		sh

.PHONY: generate
generate: _executable _generate_dir
	- docker build --target app_protoc -t app_cli .docker/php/cli
	- docker run --init -it --rm \
		-u $(USER) \
		-v "$$(pwd)/bin:/app/bin" \
		-v "$$(pwd)/vendor:/app/vendor" \
		-v "$$(pwd)/$(GENERATE_DIR):/app/$(GENERATE_DIR)" \
		-v "$$(pwd)/../protos/proto:/app/proto" \
		-w /app \
		app_cli sh -c \
		    "protoc -I proto proto/sso/*.proto --plugin=protoc-gen-custom-plugin=bin/compiler.php --custom-plugin_out=$(GENERATE_DIR)"
	- docker image rm -f app_cli

.PHONY: run
run:
	- docker build --target app_devel -t app_cli .docker/php/cli
	- docker run --init -it --rm \
		--add-host=host.docker.internal:host-gateway \
		--env-file .docker/environment/base.env \
		-u $(USER) \
		-v "$$(pwd)/src:/app/src" \
		-v "$$(pwd)/vendor:/app/vendor" \
		-v "$$(pwd)/$(GENERATE_DIR):/app/$(GENERATE_DIR)" \
		-w /app/src \
		app_cli php index.php
	- docker image rm -f app_cli

_executable:
	chmod +x bin/compiler.php

_generate_dir:
	until [ -d $(GENERATE_DIR) ] ; do mkdir -p ${GENERATE_DIR} ; done
