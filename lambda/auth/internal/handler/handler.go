package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/yunomu/later-execution/lib/db/accesskeydb"
)

type Handler struct {
	db accesskeydb.DB
}

func NewHandler(
	db accesskeydb.DB,
) *Handler {
	return &Handler{
		db: db,
	}
}

type Request events.APIGatewayV2CustomAuthorizerV2Request

func (h *Handler) Serve(ctx context.Context, event *Request) error {
	// TODO
	return nil
}
