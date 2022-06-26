package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"github.com/yunomu/later-execution/lib/db/accesskeydb"
	"github.com/yunomu/later-execution/lib/db/executiondb"
)

type Handler struct {
	accesskeydb accesskeydb.DB
	executiondb executiondb.DB
}

func NewHandler(
	accesskeydb accesskeydb.DB,
	executiondb executiondb.DB,
) *Handler {
	return &Handler{
		accesskeydb: accesskeydb,
		executiondb: executiondb,
	}
}

type Request events.APIGatewayV2HTTPRequest
type Response events.APIGatewayV2HTTPResponse

func (h *Handler) Serve(ctx context.Context, event *Request) (*Response, error) {
	// TODO
	return &Response{
		StatusCode: 503,
	}, nil
}
