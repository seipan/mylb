DOCKER_FILE_DIR                 := ./docker
DOCKER_COMPOSE_LOCAL_SERVER     := $(DOCKER_FILE_DIR)/docker-compose.yml

.PHONY: tidy
tidy: 
	./script/gomod.sh


.PHONY: up
up:
	docker-compose \
	-f $(DOCKER_COMPOSE_LOCAL_SERVER) up -d


.PHONY: down
down:
	docker-compose \
	-f $(DOCKER_COMPOSE_LOCAL_SERVER) down \
	--rmi all --volumes --remove-orphans


.PHONY: re
re:down up