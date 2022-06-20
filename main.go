package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/Carlsberg/configuration-fetch-action/aws"
	"github.com/crqra/go-action/pkg/action"
	"github.com/itchyny/gojq"
)

func main() {
	if err := action.Execute(&AppConfigFetchAction{}); err != nil {
		action.SetFailed(err, map[string]string{})
	}
}

type AppConfigFetchAction struct {
	AppName     string `action:"app-name"`
	ProfileName string `action:"profile-name"`
	Environment string `action:"env"`
	Region      string `action:"region"`
}

func (a *AppConfigFetchAction) Run() error {
	config, err := aws.GetConfig(context.Background(), a.AppName, a.ProfileName, a.Environment, a.Region)
	if err != nil {
		return err
	}

	var sb strings.Builder

	query, err := gojq.Parse(".")
	if err != nil {
		return err
	}

	iter := query.Run(config)

	for {
		v, ok := iter.Next()
		if !ok {
			break
		}

		if err, ok := v.(error); ok {
			return err
		}

		sb.WriteString(fmt.Sprintf("%#v\n", v))
	}
	action.SetOutput("config", sb.String())

	return err
}
