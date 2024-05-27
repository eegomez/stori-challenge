package file

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"io"
	"log"
)

func newAwsRepository(cfg *configuration.Config) Repository {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.S3BucketRegion))
	if err != nil {
		log.Fatal(fmt.Errorf("unable to load SDK config, %v", err))
	}
	client := s3.NewFromConfig(awsCfg)
	return &awsRepository{
		config: cfg,
		client: client,
	}
}

type awsRepository struct {
	config *configuration.Config
	client *s3.Client
}

func (repo *awsRepository) GetFile(ctx context.Context, filename string) (io.Reader, error) {
	output, err := repo.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(repo.config.S3BucketName),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get file from S3, %v", err)
	}

	return output.Body, nil
}
