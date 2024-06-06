package dynamo

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "log"
    "go-dynamodb-crud/model"
)

var (
    svc *dynamodb.DynamoDB
    tableName = "Items"
)

func InitDynamo() {
    sess, err := session.NewSession(&aws.Config{
        Region:   aws.String("us-west-2"),
        Endpoint: aws.String("http://dynamodb-local:8000"),
    })
    if err != nil {
        log.Fatalf("Failed to connect to DynamoDB: %v", err)
    }
    svc = dynamodb.New(sess)
    createTable()
}

func createTable() {
    input := &dynamodb.CreateTableInput{
        TableName: aws.String(tableName),
        KeySchema: []*dynamodb.KeySchemaElement{
            {
                AttributeName: aws.String("ID"),
                KeyType:       aws.String("HASH"),
            },
        },
        AttributeDefinitions: []*dynamodb.AttributeDefinition{
            {
                AttributeName: aws.String("ID"),
                AttributeType: aws.String("S"),
            },
        },
        ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
            ReadCapacityUnits:  aws.Int64(5),
            WriteCapacityUnits: aws.Int64(5),
        },
    }

    _, err := svc.CreateTable(input)
    if err != nil && err.Error() != dynamodb.ErrCodeResourceInUseException {
        log.Fatalf("Got error calling CreateTable: %s", err)
    }
}

func PutItem(item model.Item) error {
    av, err := dynamodbattribute.MarshalMap(item)
    if err != nil {
        return err
    }

    input := &dynamodb.PutItemInput{
        Item:      av,
        TableName: aws.String(tableName),
    }

    _, err = svc.PutItem(input)
    return err
}

func GetItem(id string) (*model.Item, error) {
    result, err := svc.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "ID": {
                S: aws.String(id),
            },
        },
    })

    if err != nil {
        return nil, err
    }

    if result.Item == nil {
        return nil, nil
    }

    item := new(model.Item)
    err = dynamodbattribute.UnmarshalMap(result.Item, item)
    if err != nil {
        return nil, err
    }

    return item, nil
}

func DeleteItem(id string) error {
    _, err := svc.DeleteItem(&dynamodb.DeleteItemInput{
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "ID": {
                S: aws.String(id),
            },
        },
    })
    return err
}

func ListItems() ([]model.Item, error) {
    result, err := svc.Scan(&dynamodb.ScanInput{
        TableName: aws.String(tableName),
    })

    if err != nil {
        return nil, err
    }

    items := []model.Item{}
    err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
    if err != nil {
        return nil, err
    }

    return items, nil
}