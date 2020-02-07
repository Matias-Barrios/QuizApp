package config

import (
	"fmt"
	"os"
)

// EnvironmentFetcher :
type EnvironmentFetcher struct {
	IenvironmentFetcher
}

// GetValue :
func (e EnvironmentFetcher) GetValue(key string) (string, error) {
	value, err := os.LookupEnv(key)
	if err == false {
		return "", fmt.Errorf("Unable to find '%s' in the environment", key)
	}
	return value, nil
}
