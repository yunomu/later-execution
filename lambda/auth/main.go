package main

import (
	"context"
	"os"

	"go.uber.org/zap"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/yunomu/later-execution/lib/db/accesskeydb"

	"github.com/yunomu/later-execution/lambda/auth/internal/handler"
)

var logger *zap.Logger

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = l
}

func main() {
	ctx := context.Background()

	region := os.Getenv("REGION")
	accessKeyTable := os.Getenv("ACCESS_KEY_TABLE")

	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		logger.Fatal("NewSession", zap.Error(err))
	}

	h := handler.NewHandler(
		accesskeydb.NewDynamoDB(
			dynamodb.New(sess),
			accessKeyTable,
		),
	)

	lambda.StartWithContext(ctx, h.Serve)
}
