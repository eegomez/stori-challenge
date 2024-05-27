package file

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"io"
	"log"
)

var ErrGetObjectFailed = errors.New("failed to get file from S3")

type S3ClientInterface interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

func newAwsFileRepositoryFactory(cfg *configuration.Config) Repository {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.S3BucketRegion))
	if err != nil {
		log.Fatal(fmt.Errorf("unable to load SDK config, %v", err))
	}
	client := s3.NewFromConfig(awsCfg)
	return newAwsFileRepository(cfg, client)
}

func newAwsFileRepository(cfg *configuration.Config, client S3ClientInterface) Repository {
	return &awsRepository{
		config: cfg,
		client: client,
	}
}

type awsRepository struct {
	config *configuration.Config
	client S3ClientInterface
}

func (repo *awsRepository) GetFile(ctx context.Context, filename string) (io.Reader, error) {
	output, err := repo.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(repo.config.S3BucketName),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, ErrGetObjectFailed
	}

	return output.Body, nil
}
