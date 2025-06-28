# Makefile para o projeto Currency Exchange Server

# Variáveis
BINARY_NAME=currency-exchange-server
BUILD_DIR=build

# Comando de ajuda
.PHONY: help
help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Comando para rodar o projeto
.PHONY: run
run: ## Executa o projeto
	@echo "Executando o projeto..."
	go run cmd/main.go

# Comando para buildar o projeto
.PHONY: build
build: ## Compila o projeto gerando um executável
	@echo "Buildando o projeto..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/main.go
	@echo "Executável gerado em: $(BUILD_DIR)/$(BINARY_NAME)"

# Comando para testar o projeto
.PHONY: test
test: ## Executa os testes do projeto
	@echo "Executando testes..."
	go test ./... 