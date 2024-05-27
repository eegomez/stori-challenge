package configuration

import (
	"log"
	"os"
)

type Config struct {
	FromEmail           string
	FromEmailPassword   string
	FileImplementation  string
	EmailImplementation string
	S3BucketName        string
	S3BucketRegion      string
}

func LoadConfig() *Config {
	config := Config{
		FromEmail:           getEnv("EMAIL_ACCOUNT", "non-existent-email-123123123421dsgsdcgsdf@gmail.com"),
		FromEmailPassword:   getEnv("EMAIL_PASSWORD", "fake-password"),
		FileImplementation:  getEnv("FILE_IMPLEMENTATION", "local"),
		EmailImplementation: getEnv("EMAIL_IMPLEMENTATION", "local"),
		S3BucketName:        getEnv("S3_BUCKET_NAME", "non-existent-bucket"),
		S3BucketRegion:      getEnv("S3_BUCKET_REGION", "us-east-1"),
	}

	if config.FromEmail == "" || config.FromEmailPassword == "" || config.FileImplementation == "" || config.EmailImplementation == "" || config.S3BucketName == "" || config.S3BucketRegion == "" {
		log.Fatal("Missing required environment variables")
	}

	return &config
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
