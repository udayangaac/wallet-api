// Package cfgloader responsible for loading configurations.
package cfgloader

import "os"

// LookupEnv check value for specific key in the environment or return value.
func LookupEnv(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}
