package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/simon-watiau/logcop/dto"
	"github.com/simon-watiau/logcop/pattern-extractor/cluster"
	"github.com/simon-watiau/logcop/pattern-extractor/config"
	"github.com/simon-watiau/logcop/subscriber"
	"github.com/simon-watiau/logcop/tokenizer"
	"go.uber.org/zap"
)

func main() {
	logger, loggerErr := zap.NewProduction()
	if loggerErr != nil {
		logger.Fatal(
			"failed to create logger",
			zap.String("error", loggerErr.Error()),
		)

		os.Exit(1)
	}

	defer logger.Sync()

	config, configErr := config.GetConfig()
	if configErr != nil {
		logger.Fatal(
			"invalid config",
			zap.String("error", configErr.Error()),
		)

		os.Exit(1)
	}

	subscriber := subscriber.NewSqsClient(subscriber.SqsSubscriberConfig{
		AwsSqsEndpoint:          config.AwsSqsEndpoint,
		AwsSqsRegion:            config.AwsSqsRegion,
		AwsAccessKeyId:          config.AwsAccessKeyId,
		AwsSecretAccessKey:      config.AwsSecretAccessKey,
		AwsSqsExtractorQueueUrl: config.AwsSqsExtractorQueueUrl,
		AwsSqsSigninRegion:      config.AwsSqsSigninRegion,
	})

	subscriber.InitClient()

	for {
		message, receipt, err := subscriber.Receive()

		if err != nil {
			logger.Error(
				"failed to read message",
				zap.String("error", err.Error()),
				zap.String("queue", config.AwsSqsExtractorQueueUrl),
			)
			continue
		}

		if message == nil {
			continue
		}

		var rawMessage dto.BatchDto

		unmarshallErr := json.Unmarshal(
			[]byte(*message),
			&rawMessage,
		)

		if unmarshallErr != nil {
			logger.Info(
				"Failed to unmarshal",
				zap.Error(unmarshallErr),
			)
			continue
		}

		fullPath := fmt.Sprintf(
			"%s%c%s",
			config.BatchDir,
			os.PathSeparator,
			rawMessage.FileName,
		)
		file, err := os.Open(fullPath)
		if err != nil {
			logger.Info(
				"Failed to read file",
			)
			continue
		}

		scanner := bufio.NewScanner(file)
		clusterSet := cluster.NewClusterAggregate(
			context.Background(),
			[]float64{
				0.01,
				0.1,
				0.3,
			})
		for scanner.Scan() {
			var log dto.LogDto
			json.Unmarshal([]byte(scanner.Text()), &log)
			var m map[string]string
			json.Unmarshal([]byte(log.Body), &m)

			clusterSet.AddLog(tokenizer.NewFromString(m["short_message"]))
		}

		file.Close()

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		for _, e := range clusterSet.Aggregate() {

			fmt.Printf("%s\n", e)

		}

		subscriber.Ack(receipt)
	}
}
