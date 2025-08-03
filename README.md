# Currency Exchange

Este projeto é uma aplicação para consulta da cotação do dólar, composta por dois serviços principais: um servidor (API) e um cliente de linha de comando. O servidor integra com fontes externas de cotação e armazena os dados em um banco de dados. O cliente permite consultar a cotação do dólar de forma interativa.

## Estrutura do Projeto

- `server/`: Serviço de API, responsável por buscar e armazenar a cotação do dólar.
- `client/`: Cliente CLI, que consome a API para exibir a cotação ao usuário.
- `Makefile`: Arquivo de automação na raiz para facilitar o uso dos serviços.

## Pré-requisitos
- Docker e Docker Compose
- Go 1.21 ou superior
- Make

## Como executar

1. **Subir o banco de dados**

```bash
make start-db
```

2. **Executar o servidor (API)**

```bash
make start-server
```

3. **Executar o cliente**

```bash
make start-client
```

4. **Executar tudo de uma vez**

```bash
make run
```

## Comandos disponíveis
Consulte todos os comandos possíveis com:

```bash
make help
```

## Funcionamento

- O servidor busca a cotação do dólar em provedores externos e salva no banco de dados.
- O cliente permite ao usuário consultar a cotação do dólar digitando comandos interativos.

## Exemplos de uso do cliente

Ao rodar o cliente, digite:
- `dollar` para consultar a cotação do dólar
- `exit` para sair

## Licença
Este projeto é open source e está sob a licença MIT.