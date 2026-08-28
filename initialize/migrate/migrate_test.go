package migrate

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/perfect-panel/server/pkg/orm"
)

func TestMigrateMySQL(t *testing.T) {
	dsn := os.Getenv("PPANEL_TEST_MYSQL_DSN")
	if dsn == "" {
		t.Skip("set PPANEL_TEST_MYSQL_DSN to run MySQL/MariaDB migration test")
	}
	runMigration(t, orm.DriverMySQL, dsn)
}

func TestMigratePostgres(t *testing.T) {
	dsn := os.Getenv("PPANEL_TEST_POSTGRES_DSN")
	if dsn == "" {
		t.Skip("set PPANEL_TEST_POSTGRES_DSN to run PostgreSQL migration test")
	}
	runMigration(t, orm.DriverPostgres, dsn)
}

func runMigration(t *testing.T, driver, dsn string) {
	t.Helper()
	err := Up(driver, dsn)
	if err != nil && !errors.Is(err, NoChange) {
		t.Fatalf("%s migration failed: %v", driver, err)
	}
	cfg := orm.ParseDSN(dsn)
	if cfg == nil {
		t.Fatalf("%s dsn parse failed", driver)
	}
	cfg.Driver = orm.NormalizeDriver(driver)
	db, err := orm.ConnectDatabase(orm.Mysql{Config: *cfg})
	if err != nil {
		t.Fatalf("%s connect failed: %v", driver, err)
	}
	sqlDB, err := db.DB()
	if err == nil {
		defer sqlDB.Close()
	}
	if err := CreateAdminUser(fmt.Sprintf("admin-%s@example.com", driver), "password", db); err != nil {
		t.Fatalf("%s create admin failed: %v", driver, err)
	}
	if err := CreateAdminUser("", "password", db); err != nil {
		t.Fatalf("%s existing admin must skip email validation: %v", driver, err)
	}
}

type fakeMigrationRunner struct {
	upErr       error
	sourceErr   error
	databaseErr error
	closed      bool
}

func (f *fakeMigrationRunner) Up() error { return f.upErr }

func (f *fakeMigrationRunner) Close() (error, error) {
	f.closed = true
	return f.sourceErr, f.databaseErr
}

func TestUpAndCloseAlwaysClosesMigrationRunner(t *testing.T) {
	upFailure := errors.New("up failed")
	tests := []struct {
		name  string
		upErr error
	}{
		{name: "success"},
		{name: "no change", upErr: NoChange},
		{name: "migration failure", upErr: upFailure},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := &fakeMigrationRunner{upErr: tt.upErr}
			err := upAndClose(runner)

			if !runner.closed {
				t.Fatal("migration runner was not closed")
			}
			if !errors.Is(err, tt.upErr) {
				t.Fatalf("upAndClose() error = %v, want %v", err, tt.upErr)
			}
		})
	}
}

func TestUpAndCloseDoesNotHideCloseFailureBehindNoChange(t *testing.T) {
	closeFailure := errors.New("database close failed")
	runner := &fakeMigrationRunner{
		upErr:       NoChange,
		databaseErr: closeFailure,
	}

	err := upAndClose(runner)

	if !runner.closed {
		t.Fatal("migration runner was not closed")
	}
	if !errors.Is(err, closeFailure) {
		t.Fatalf("upAndClose() error = %v, want close failure", err)
	}
	if errors.Is(err, NoChange) {
		t.Fatalf("upAndClose() error = %v hides a close failure behind ErrNoChange", err)
	}
}
