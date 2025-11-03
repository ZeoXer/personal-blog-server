package r2client

import (
	"context"
	"go-server/global"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	S3 *s3.Client
}

func NewR2Client(ctx context.Context) (*Client, error) {
	endpoint := global.CONFIG.R2Storage.Endpoint
	region := global.CONFIG.R2Storage.Region
	accessKey := global.CONFIG.R2Storage.AccessKey
	secretKey := global.CONFIG.R2Storage.SecretKey

	s3c := s3.New(s3.Options{
		Region:       region,
		Credentials:  aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		UsePathStyle: true,
		BaseEndpoint: aws.String(endpoint),
	})

	return &Client{
		S3: s3c,
	}, nil
}
