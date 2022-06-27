package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Carlsberg/configuration-fetch-action/aws"
)

func main() {
	appName := os.Getenv("INPUT_APP_NAME")
	profileName := os.Getenv("INPUT_PROFILE_NAME")
	env := os.Getenv("INPUT_ENV")
	region := os.Getenv("INPUT_REGION")

	config, err := aws.GetConfig(context.Background(), appName, profileName, env, region)
	if err != nil {
		log.Fatal(err)
	}

	config = strings.Replace(config, "\n", "", -1)

	fmt.Println(fmt.Sprintf("::set-output name=config::%v", config))
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

	config = strings.Replace(config, "\n", "", -1)

	return err
}
