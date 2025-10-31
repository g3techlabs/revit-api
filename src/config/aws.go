package config

import (
	"context"
	"log"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client() *s3.Client {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("unable to load SDK config: " + err.Error())
	}

	usePathStyle := Get("AWS_ENDPOINT_URL") != ""
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = usePathStyle
	})
}

func NewPresignClient(s3Client *s3.Client) *s3.PresignClient {
	return s3.NewPresignClient(s3Client)
}
