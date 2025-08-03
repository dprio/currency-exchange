# Makefile na raiz do projeto currency-exchange


# Comando de ajuda
.PHONY: help
help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


# Comandos para o servidor (API)
.PHONY: start-server
start-server: ## Executa o servidor (API) na pasta server
	cd server &&  docker-compose up -d && make -s run &

# Comandos para o cliente
.PHONY: start-client
start-client: ## Executa o cliente na pasta client
	cd client && make run

.PHONY: run
run: ## Sobe o server e client
	@echo "Iniciando o server..."
	make -s start-server
	sleep 2
	@echo "Iniciando o client..."
	make -s start-client
	make -s clean

.PHONY: clean
clean: ## Limpa os containers do docker
	@echo "Limpando..."
	cd server && docker-compose down
