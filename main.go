package main

import (
	"github.com/Carlsberg/configuration-fetch-action/aws"
	"github.com/crqra/go-action/pkg/action"
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
	config, err := aws.GetConfig(a.AppName, a.ProfileName, a.Environment, a.Region)

	action.SetOutput("config", config)

	return err
}
