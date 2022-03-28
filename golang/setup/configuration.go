package setup

import (
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	FEED_API_KEY                 string
	REPOSITORY_CONNECTION_STRING string
	PASSWORD_HASH_METHOD         string
}

var configuration *Configuration

func init() {
	configuration = &Configuration{}
	var err error

	defer func() {
		if err != nil {
			log.Panicf("CONFIGURATION ERROR %v", err)
		}
	}()
	if configuration.FEED_API_KEY, err = requiredEnv("FEED_API_KEY"); err != nil {
		return
	}
	if configuration.REPOSITORY_CONNECTION_STRING, err = requiredEnv("REPOSITORY_CONNECTION_STRING"); err != nil {
		return
	}
	configuration.PASSWORD_HASH_METHOD = optionalEnv("PASSWORD_HASH_METHOD", "sha256")
}

func requiredEnv(envName string) (value string, err error) {
	value = os.Getenv(envName)
	if len(value) == 0 {
		return value, fmt.Errorf("MISSING ENVIRONMENT '%s'", envName)
	}
	return value, nil
}

func optionalEnv(envName string, defaultValue string) (value string) {
	value = os.Getenv(envName)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func GetConfiguration() *Configuration {
	return configuration
}
