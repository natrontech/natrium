package server

import (
	"github.com/natrongmbh/natrium/pkg/env"
	"github.com/natrongmbh/natrium/pkg/migrations"
	"github.com/pocketbase/pocketbase"
)

var App *pocketbase.PocketBase

// Setup initializes the pocketbase server
func Setup() error {

	// initialize pocketbase collections
	migrations.InitCollections()

	// initialize pocketbase server
	App = pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir:       env.POCKETBASE_DATA_DIR,
		DefaultEncryptionEnv: env.POCKETBASE_ENCRYPTION_KEY,
	})

	return nil
}
