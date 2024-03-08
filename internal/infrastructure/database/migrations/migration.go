package migrations

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RunMigrations executes all migration scripts found in the predefined directory.
func RunMigrations(db *sql.DB, migrationsDir string) error {

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migration directory: %w", err)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			file := filepath.Clean(file.Name())
			if !strings.HasPrefix(file, "migration") {
				err = fmt.Errorf("migration script %s does not follow naming convention", file)
				return errors.New(err.Error())
			}

			migrationPath := filepath.Join(migrationsDir, file)

			migrationScript, err := os.ReadFile(filepath.Clean(migrationPath))
			if err != nil {
				return fmt.Errorf("failed to read migration script %s: %w", file, err)
			}

			_, err = db.Exec(string(migrationScript))
			if err != nil {
				return fmt.Errorf("failed to execute migration script %s: %w", file, err)
			}
		}
	}

	return nil
}
