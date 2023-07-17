package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"os"
)

const migrationPathEnv = "MIGRATION_PATH"

func ApplySQLScript(db *sqlx.DB) error {
	filePath := os.Getenv(migrationPathEnv)

	scriptBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL script file: %w", err)
	}

	_, err = db.Exec(string(scriptBytes))
	if err != nil {
		return fmt.Errorf("failed to apply SQL script: %w", err)
	}

	return nil
}
