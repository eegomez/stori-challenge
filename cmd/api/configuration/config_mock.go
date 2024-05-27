package configuration

func MockDefaultConfig() *Config {
	return &Config{
		FromEmail:           "some.email@gmail.com",
		FromEmailPassword:   "some-password",
		FileImplementation:  "some-file-implementation",
		EmailImplementation: "some-email-implementation",
		S3BucketName:        "some-bucket-name",
		S3BucketRegion:      "some-bucket-region",
	}
}
