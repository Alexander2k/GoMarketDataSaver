package postgresStorage

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log/slog"
)

func (d *PostgresDB) Migrate() error {
	driver, err := postgres.WithInstance(d.Db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://app/migrations/postgres", "users", driver)
	if err != nil {
		slog.Error("Could find migrations: %v", err)
		return err
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("cant migrate: %v", err)
		}

	}

	slog.Info("migrations completed")
	return nil
}
