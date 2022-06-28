package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

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
	config, err := aws.GetConfig(context.Background(), a.AppName, a.ProfileName, a.Environment, a.Region)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%v-%v-%v-%v.json", a.Region, a.AppName, a.Environment, a.ProfileName)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(config)
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	action.SetOutput("path", absPath)

	return err
}
