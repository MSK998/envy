package envy

import "os"

// Will look for the environment variable, if unavailable,
// returning the fallback as a string.
func GetEnvOrDefault(envVar string, fallback string) string {
	if value, exists := os.LookupEnv(envVar); exists {
		return value
	}
	return fallback
}
