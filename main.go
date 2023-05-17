package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/konjure-examples/go-aws-deployment/internal/awswrapper"
	"github.com/konjure-examples/go-aws-deployment/internal/handler"
)

const streamNameEnvKey = "KINESIS_STREAM_NAME"

const dynamodbTableEnvKey = "DYNAMODB_TABLE"

func main() {
	streamName := os.Getenv(streamNameEnvKey)
	dynamodbTableName := os.Getenv(dynamodbTableEnvKey)
	bucketName := os.Getenv("S3" + "_BUCKET")
	queueUrl := os.Getenv("SQS_URL")
	ctx := context.Background()
	wrapper, err := awswrapper.New(ctx, streamName, dynamodbTableName, bucketName, queueUrl)
	if err != nil {
		return
	}

	h := handler.NewHandler(wrapper)

	if err = http.ListenAndServe(os.Getenv("PORT"), h); err != nil {
		log.Fatal(err)
	}
}
