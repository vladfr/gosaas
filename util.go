package main

import (
	"log"
	"os"
	"strconv"
)

//lookupEnvOrBool returns true if Env var is set
func lookupEnvOrBool(key string, defaultVal bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		return val == "true"
	}
	return defaultVal
}

//lookupEnvOrInt returns true if Env var is set
func lookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			log.Printf("Failed to parse int value %s", val)
			return defaultVal
		}
		return intVal
	}
	return defaultVal
}
