package config

import (
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

// Server
var ProductionMode = getEnvOrSetFallbackBool("APP_PRODUCTION_MODE", false)
var EnableDetailedRequestLogging = getEnvOrSetFallbackBool("APP_DETAILED_REQUEST_LOGGING", true)

var EnableTls = getEnvOrSetFallbackBool("APP_ENABLE_TLS", true)
var GenerateSelfSignedCert = getEnvOrSetFallbackBool("APP_TLS_GENERATE_SELF_SIGNED_CERTs", true)
var CertValidityDays = getEnvOrSetFallbackInt("APP_SELF_SIGNED_CERT_VALIDITY_DAYS", 30)
var CertPath = getEnvOrSetFallbackString("APP_TLS_CERT_PATH", "./cert.pem")
var KeyPath = getEnvOrSetFallbackString("APP_TLS_KEY_PATH", "./key.pem")

// Comma separated list of hosts and IPs for self-signed cert
var TlsCertificateHosts = getEnvOrSetFallbackString("APP_TLS_CERT_HOSTS", "localhost")

var ServerAddress = getEnvOrSetFallbackString("APP_SRV_ADDRESS", "0.0.0.0:8080")
var ServerReadHeaderTimeout = getEnvOrSetFallbackInt("APP_SRV_READ_HEADER_TIMEOUT", 10)
var ServerReadTimeout = getEnvOrSetFallbackInt("APP_SRV_READ_TIMEOUT", 20)
var ServerWriteTimeout = getEnvOrSetFallbackInt("APP_SRV_WRITE_TIMEOUT", 20)
var ServerIdleTimeout = getEnvOrSetFallbackInt("APP_SRV_IDLE_TIMEOUT", 45)

// Gets value from environment variable or sets fallback (default) value. Returns string. Case sensitive.
func getEnvOrSetFallbackString(envVar string, fallbackValue string) string {
	val := os.Getenv(envVar)
  if val == "" {
    val = fallbackValue
  }

  return val
}

// Gets value from environment variable or sets fallback (default) value. Returns parsed value as bool. Case sensitive.
func getEnvOrSetFallbackBool(envVar string, fallbackValue bool) bool {
	val := os.Getenv(envVar)
	if val == "" {
		return fallbackValue
	}

	valBool, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatal().Err(err).Str("environment-variable", envVar).Msg("Error when parsing configuration")
	}

	return valBool
}

// Gets value from environment variable or sets fallback (default) value. Returns parsed value as int. Case sensitive.
func getEnvOrSetFallbackInt(envVar string, fallbackValue int) int {
	val := os.Getenv(envVar)
	if val == "" {
		return fallbackValue
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal().Err(err).Str("environment-variable", envVar).Msg("Error when parsing configuration")
	}

	return valInt
}
