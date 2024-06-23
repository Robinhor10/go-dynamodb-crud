#!/bin/bash

# Espera o LocalStack estar pronto
echo "Waiting for LocalStack to be ready..."
TIMEOUT=1
WAIT_INTERVAL=5
ELAPSED=0

# Verificar se o LocalStack está pronto com base na URL de saúde
until curl -s http://localhost:4566/health | grep "\"dynamodb\": \"running\"" > /dev/null; do
    echo "Waiting for LocalStack to be ready... ($ELAPSED/$TIMEOUT seconds)"
    curl -s http://localhost:4566/health # Adicionando este log para diagnóstico
    sleep $WAIT_INTERVAL
    ELAPSED=$((ELAPSED + WAIT_INTERVAL))
    if [ "$ELAPSED" -ge "$TIMEOUT" ]; then
        echo "Timeout waiting for LocalStack to be ready"
        exit 1
    fi
done

# Executa o script de criação da tabela
/docker-entrypoint-initaws.d/create_table.sh

# Executa o entrypoint padrão do LocalStack
exec docker-entrypoint.sh
