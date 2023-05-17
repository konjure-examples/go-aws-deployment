package handler

import (
	"encoding/json"
	"net/http"

	"github.com/konjure-examples/go-aws-deployment/internal/awswrapper"
)

type Handler struct {
	awsWrapper *awswrapper.AWSWrapper
}

func NewHandler(awsWrapper *awswrapper.AWSWrapper) *Handler {
	return &Handler{awsWrapper: awsWrapper}
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	switch request.URL.Path {
	case "s3":
		var s3Object awswrapper.S3Object
		err := json.NewDecoder(request.Body).Decode(&s3Object)
		if err != nil {
			return
		}

		err = h.awsWrapper.PutObjectWrapper(ctx, request.URL.Query().Get("key"), &s3Object)
		if err != nil {
			return
		}
		writer.WriteHeader(200)

	case "dynamodb":
		key := request.URL.Query().Get("key")
		prefix := request.URL.Query().Get("prefix")
		err := h.awsWrapper.QueryTableWrapper(ctx, key, prefix)
		if err != nil {
			return
		}

		err = h.awsWrapper.GetItemWrapper(ctx, key)
		if err != nil {
			return
		}

	case "sqs":
		err := h.awsWrapper.ReceiveMessageWrapper(ctx)
		if err != nil {
			return
		}

	case "kinesis":
		err := h.awsWrapper.ListShardsWrapper(ctx)
		if err != nil {
			return
		}

		var record awswrapper.KinesisRecord
		err = json.NewDecoder(request.Body).Decode(&record)

		err = h.awsWrapper.PutKinesisRecordWrapper(ctx, &record)
		if err != nil {
			return
		}
	}

}
