DOCKER_COMPOSE_LOCAL_SERVER     := docker-compose.yml

.PHONY: ts-re
ts-re: 
	cd testserver && make re

.PHONY: ts-down
ts-down: 
	cd testserver && make down

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