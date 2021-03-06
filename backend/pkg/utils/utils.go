package utils

import (
	"os"
	"strconv"
)

func GetEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func GetPortEnv(key string, defaultVal uint16) uint16 {
	if val := os.Getenv(key); val != "" {
		integer, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			panic("Error to get Postgres port environment")
		}
		return uint16(integer)
	}
	return defaultVal
}

func GetIntEnv(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		integer, err := strconv.Atoi(val)
		if err != nil {
			panic("Error to get Postgres port environment")
		}
		return integer
	}
	return defaultVal
}

// CheckIfEnvironmentExists return true/false if the environment variable exists
func CheckIfEnvironmentExists(env string) bool {
	_, exist := os.LookupEnv(env)
	return exist
}
