# Go DynamoDB CRUD API

Esta aplicação é uma API RESTful escrita em Go que realiza operações CRUD em uma tabela DynamoDB. 

## Endpoints

- `POST /items`: Cria um novo item.
- `GET /items/{id}`: Retorna um item pelo ID.

## Executando a aplicação

### Pré-requisitos

- Docker
- Docker Compose

### Passos

1. Clone o repositório:

    git clone `https://github.com/Robinhor10/go-dynamodb-crud.git`
    Digite no terminal dentro da IDE:cd myapp

2. Suba os containers:

   Digite no terminal dentro da IDE: `docker-compose up --build`

Isso vai iniciar a aplicação na porta 8080 e o DynamoDB Local na porta 8000.

3. Acesse o container localstack 

    3.1 Execute a configuração da aws digitando `aws configure`, após isso configure com as informações abaixo.
    `
    AWS Access Key ID [None]: test
    AWS Secret Access Key [None]: test
    Default region name [None]: us-east-1
    Default output format [None]:
    `
    Digite ENTER

    3.2 Execute o código abaixo para que a tabela seja criada
    `
    aws dynamodb create-table \
    --table-name Items \
    --attribute-definitions \
        AttributeName=id,AttributeType=S \
    --key-schema \
        AttributeName=id,KeyType=HASH \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --region us-east-1 \
    --endpoint-url=http://localhost:4566
    `

### Exemplos de uso

#### Criar um item

curl -X POST -H "Content-Type: application/json" -d '{"id":"1","name":"Item 1"}' http://localhost:8080/items

#### Criar um item

curl http://localhost:8080/items/1