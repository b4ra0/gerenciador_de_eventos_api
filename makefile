# makefile para build e run de docker-compose

# Variáveis
DOCKER_COMPOSE_FILE = docker-compose.yaml

# Alvos
.PHONY: all build run

all: build run

build:
	@echo "Construindo containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

run:
	@echo "Iniciando containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

clean:
	@echo "Parando e removendo containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down