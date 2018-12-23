package cmd

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var env *environment

type environment struct {
	GcpProjectID string `envconfig:"GCP_PROJECT_ID" default:""`
}

func loadEnvironmentVariables() error {
	e := &environment{}
	if err := envconfig.Process("", e); err != nil {
		return fmt.Errorf("failed to load environment variables: %s", err)
	}

	env = e
	return nil
}
