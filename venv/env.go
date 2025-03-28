package venv

import "os"

const (
	AppEnvKey  = "APP_ENV"
	AppEnvDev  = "dev"
	AppEnvTest = "test"
	AppEnvPre  = "pre"
	AppEnvProd = "prod"
)

// GetEnv get env, support default value
func GetEnv(key string, defaultVal string) (value string) {
	value = os.Getenv(key)
	if value != "" {
		return
	}

	return defaultVal
}
