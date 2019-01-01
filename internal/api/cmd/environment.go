package cmd

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var env *environment

type environment struct {
	GcpProjectID           string `envconfig:"GCP_PROJECT_ID" default:""`
	EnableStackdriverTrace bool   `envconfig:"ENABLE_STACKDRIVER_TRACE" default:"false"`
	AppEngineNamespace     string `envconfig:"APP_ENGINE_NAMESPACE" default:""`
}

func loadEnvironmentVariables() error {
	e := &environment{}
	if err := envconfig.Process("", e); err != nil {
		return fmt.Errorf("failed to load environment variables: %s", err)
	}

	env = e
	return nil
}
