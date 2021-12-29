package core

import (
	"log"
	"strconv"
)

const (
	port        = "PORT"
	defaultPort = 3000
)

func LoadConfig(getter getFn) (Config, error) {
	return Config{
		Port: getInt(getter, port, defaultPort),
	}, nil
}

func getInt(getter getFn, key string, defaultValue int) int {
	raw := getter(key)

	if raw == "" {
		log.Printf("No variable %s found, using default %d", key, defaultValue)

		return defaultValue
	}

	result, err := strconv.ParseInt(raw, 10, 16)
	if err != nil {
		log.Printf("Error parsing value %s, using default %d", raw, defaultValue)

		return defaultValue
	}

	return int(result)
}
