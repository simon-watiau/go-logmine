package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	AwsSqsEndpoint          string `required:"true" split_words:"true"`
	AwsSqsRegion            string `required:"true" split_words:"true"`
	AwsAccessKeyId          string `required:"true" split_words:"true"`
	AwsSecretAccessKey      string `required:"true" split_words:"true"`
	AwsSqsExtractorQueueUrl string `required:"true" split_words:"true"`
	AwsSqsSigninRegion      string `required:"true" split_words:"true"`
	BatchDir                string `required:"true" split_words:"true"`
}

func GetConfig() (Config, error) {
	var config Config
	err := envconfig.Process("config", &config)
	return config, err
}
