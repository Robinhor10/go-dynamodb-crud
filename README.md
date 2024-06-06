# Go DynamoDB CRUD API

Esta aplicação é uma API RESTful escrita em Go que realiza operações CRUD em uma tabela DynamoDB. 

## Endpoints

- `POST /items`: Cria um novo item.
- `GET /items/{id}`: Retorna um item pelo ID.
- `DELETE /items/{id}`: Deleta um item pelo ID.
- `GET /items`: Lista todos os itens.

## Executando a aplicação

### Pré-requisitos

- Docker
- Docker Compose

### Passos

1. Clone o repositório:
    ```sh
    git clone <URL_DO_REPOSITORIO>
    cd go-dynamodb-crud
    ```

2. Suba os containers:
    ```sh
    docker-compose up --build
    ```

Isso vai iniciar a aplicação na porta 8080 e o DynamoDB Local na porta 8000.

### Exemplos de uso

#### Criar um item

```sh
curl -X POST http://localhost:8080/items -H "Content-Type: application/json" -d '{"id": "1", "name": "Item 1", "description": "Description 1"}'
