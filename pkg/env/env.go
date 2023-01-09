package env

import (
	"os"
)

var (
	// POCKETBASE_DATA_DIR is the default data directory
	POCKETBASE_DATA_DIR string
	// POCKETBASE_ENCRYPTION_KEY is the default encryption key
	POCKETBASE_ENCRYPTION_KEY string
	// CORS_ORIGIN is the default CORS origin
	CORS_ORIGIN string
)

// Init initializes the environment variables
func Init() error {
	if POCKETBASE_DATA_DIR = os.Getenv("POCKETBASE_DATA_DIR"); POCKETBASE_DATA_DIR == "" {
		POCKETBASE_DATA_DIR = "/pb_data"
	}

	if POCKETBASE_ENCRYPTION_KEY = os.Getenv("POCKETBASE_ENCRYPTION_KEY"); POCKETBASE_ENCRYPTION_KEY == "" {
		POCKETBASE_ENCRYPTION_KEY = "POCKETBASE_ENCRYPTION_KEY"
	}

	if CORS_ORIGIN = os.Getenv("CORS_ORIGIN"); CORS_ORIGIN == "" {
		CORS_ORIGIN = "*"
	}

	return nil
}
