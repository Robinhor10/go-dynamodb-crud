#!/bin/bash

echo "Starting table creation script..."

# Espera o LocalStack estar pronto
echo "Creating DynamoDB table..."
awslocal dynamodb create-table \
    --table-name Items \
    --attribute-definitions AttributeName=id,AttributeType=S \
    --key-schema AttributeName=id,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

echo "Table created."
