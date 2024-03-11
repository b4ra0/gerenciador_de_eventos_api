# makefile para build e run de docker-compose

# Variáveis
PORT=8080
CONTAINER_NAME=goapp
IMAGE_NAME=goapp

# Alvos
.PHONY: all build run

all: build run

build:
	@echo "Construindo containers..."
	docker build -t $(IMAGE_NAME) .

run:
	@echo "Iniciando containers..."
	docker run -d -p $(PORT):$(PORT) --name $(CONTAINER_NAME) $(IMAGE_NAME)

clean:
	@echo "Parando e removendo containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down