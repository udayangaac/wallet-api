// Package cfgloader responsible for loading configurations.
package cfgloader

import "os"

// LookupEnv get the value for specific key in the environment
// or return given default value.
func LookupEnv(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}
