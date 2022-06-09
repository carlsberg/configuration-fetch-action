package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appconfigdata"
)

func GetConfig(ctx context.Context, appName, profileName, env, region string) (string, error) {
	cfg, err := awsSdkConfig(ctx, region)
	if err != nil {
		return "", err
	}

	appconfigClient := appconfigdata.NewFromConfig(cfg)

	params := &appconfigdata.StartConfigurationSessionInput{
		ApplicationIdentifier:          &appName,
		ConfigurationProfileIdentifier: &profileName,
		EnvironmentIdentifier:          &env,
	}

	session, err := appconfigClient.StartConfigurationSession(ctx, params)
	if err != nil {
		return "", fmt.Errorf("unable to start appconfig session: %v", err)
	}

	input := &appconfigdata.GetLatestConfigurationInput{
		ConfigurationToken: session.InitialConfigurationToken,
	}

	output, err := appconfigClient.GetLatestConfiguration(ctx, input)
	if err != nil {
		return "", fmt.Errorf("unable to fetch latest configuration: %v", err)
	}

	return string(output.Configuration), nil
}

func awsSdkConfig(ctx context.Context, region string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return aws.Config{}, fmt.Errorf("unable to load SDK config, %v", err)
	}

	return cfg, nil
}
