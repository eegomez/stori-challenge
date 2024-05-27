package file

import (
	"bytes"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func Test_AWSRepository_GetFile_Successful(t *testing.T) {
	// Given
	repo, dependencies := buildFileRepository(t)
	ctx := context.Background()
	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String("some-bucket-name"),
		Key:    aws.String("some-file-name.sth"),
	}
	expectedObject := s3.GetObjectOutput{Body: io.NopCloser(bytes.NewReader([]byte("mock content")))}
	dependencies.s3client.EXPECT().GetObject(ctx, getObjectInput).
		Return(
			&expectedObject,
			nil,
		)

	// When
	reader, err := repo.GetFile(ctx, "some-file-name.sth")

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, reader)
	assert.Equal(t, reader, expectedObject.Body)
}

func Test_AWSRepository_GetFile_ErrorGettingObject(t *testing.T) {
	// Given
	repo, dependencies := buildFileRepository(t)
	ctx := context.Background()
	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String("some-bucket-name"),
		Key:    aws.String("some-file-name.sth"),
	}
	expectedError := errors.New("some-error")
	dependencies.s3client.EXPECT().GetObject(ctx, getObjectInput).Return(nil, expectedError)

	// When
	reader, err := repo.GetFile(ctx, "some-file-name.sth")

	// Then
	assert.Nil(t, reader)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrGetObjectFailed)
}

type repositoryDependencies struct {
	controller *gomock.Controller
	config     *configuration.Config
	s3client   *MockS3ClientInterface
}

func buildFileRepository(t *testing.T) (Repository, *repositoryDependencies) {
	d := buildFileRepositoryDependencies(t)
	return newAwsFileRepository(d.config, d.s3client), d
}
func buildFileRepositoryDependencies(t *testing.T) *repositoryDependencies {
	controller := gomock.NewController(t)
	return &repositoryDependencies{
		controller: controller,
		config:     configuration.MockDefaultConfig(),
		s3client:   NewMockS3ClientInterface(controller),
	}
}
