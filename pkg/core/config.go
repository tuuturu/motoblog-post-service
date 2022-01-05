package core

import (
	"log"
	"strconv"
	"strings"
)

const (
	portKey       = "PORT"
	legalHostsKey = "LEGAL_HOSTS"
	defaultPort   = 3000
	delimiter     = ";"
)

func LoadConfig(getter getFn) (Config, error) {
	return Config{
		Port:       getInt(getter, portKey, defaultPort),
		LegalHosts: getStringArray(getter, legalHostsKey, []string{}),
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

func getStringArray(getter getFn, key string, defaultValue []string) []string {
	raw := getter(key)

	if raw == "" {
		log.Printf("No variable %s found, using default %+v", key, defaultValue)

		return defaultValue
	}

	raw = strings.Trim(raw, delimiter)
	raw = strings.ReplaceAll(raw, " ", "")
	raw = strings.ReplaceAll(raw, "\n", "")

	return strings.Split(raw, delimiter)
}
