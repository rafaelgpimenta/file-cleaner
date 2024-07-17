package config

import (
	"encoding/json"
	"file-cleaner/internal/lib/logger"
	"fmt"
	"os"
	"strconv"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func init() {
	// config.ParseEnv: will parse env var in string value. eg: shell: ${SHELL}
	config.WithOptions(config.ParseEnv)

	// Add driver for support yaml content
	config.AddDriver(yaml.Driver)

	const CONFIG_DIR string = "configs"

	// Load default config
	err := config.LoadFiles(fmt.Sprintf("%s/default.yaml", CONFIG_DIR))
	if err != nil {
		logger.Fatal().Err(err).Msg("Error reading default config file")
	}

	// Get current environment (e.g., development, production)
	env := os.Getenv("APP_ENV")

	// Load env config
	if env != "" {
		err := config.LoadFiles(fmt.Sprintf("%s/%s.yaml", CONFIG_DIR, env))
		if err != nil {
			logger.Warn().Str("env", env).Err(err).Msg("Error reading config file")
		}
	}

	// Process custom env variables
	processEnvOverrides("", config.Data())
}

func processEnvOverrides(prefix string, data map[string]interface{}) {
	for key, value := range data {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}
		// Type assertions: https://go.dev/tour/methods/15
		subMap, ok := value.(map[string]interface{})
		if ok {
			processEnvOverride(fullKey, subMap)
			processEnvOverrides(fullKey, subMap)
		}
	}
}

func processEnvOverride(key string, subMap map[string]interface{}) {
	const NAME_KEY string = "__name"
	const VALUE_KEY string = "__value"

	envName, exists := subMap[NAME_KEY]
	if !exists {
		return
	}

	envValue, exists := os.LookupEnv(envName.(string))
	defaultValue := subMap[VALUE_KEY]

	if !exists {
		config.Set(key, defaultValue)
		return
	}

	var parsedValue interface{}
	var err error

	switch defaultValue.(type) {
	case string:
		parsedValue = envValue
	case int:
		parsedValue, err = strconv.Atoi(envValue)
	case float32:
		parsedValue, err = strconv.ParseFloat(envValue, 32)
	case bool:
		parsedValue, err = strconv.ParseBool(envValue)
	default:
		err = json.Unmarshal([]byte(envValue), &parsedValue)
	}

	if err != nil {
		logger.Error().Err(err).Str("key", key).Msg("Failed to parse environment variable")
		config.Set(key, defaultValue)
	} else {
		config.Set(key, parsedValue)
	}
}

func Get[K any](key string) K {
	var desiredConfig K
	config.BindStruct(key, &desiredConfig)
	return desiredConfig
}
