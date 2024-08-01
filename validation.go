package godaemon

import (
	"errors"
	"regexp"
)

const (
	DefaultEnvVariable = "DAEMON"

	PatternServiceName = "^[A-Za-z0-9_]+$"
	PatternEnvVariable = "^[A-Za-z0-9_]+$"
)

func validateServiceName(serviceName string) error {
	ok, _ := regexp.MatchString(PatternServiceName, serviceName)
	if !ok {
		return errors.New("environment variable should be consist of letters, numbers and underscores")
	}

	return nil
}

func validateEnvVariable(envVariable string) (string, error) {
	if envVariable == "" {
		return DefaultEnvVariable, nil
	}

	ok, _ := regexp.MatchString(PatternEnvVariable, envVariable)
	if !ok {
		return "", errors.New("environment variable should be consist of letters, numbers and underscores")
	}

	return envVariable, nil
}
