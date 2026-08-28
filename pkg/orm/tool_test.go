package orm

import (
	"net/url"
	"os"
	"testing"
)

func TestParseMySQLDSN(t *testing.T) {
	cfg := ParseDSN("root:password@tcp(localhost:3306)/ppanel?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	if cfg == nil {
		t.Fatal("config is nil")
	}
	if cfg.Driver != DriverMySQL || cfg.Addr != "localhost:3306" || cfg.Dbname != "ppanel" || cfg.Username != "root" {
		t.Fatalf("unexpected config: %+v", cfg)
	}
}

func TestParsePostgresDSN(t *testing.T) {
	cfg := ParseDSN("postgres://postgres:password@localhost:5432/ppanel?sslmode=disable&TimeZone=Asia%2FShanghai")
	if cfg == nil {
		t.Fatal("config is nil")
	}
	if cfg.Driver != DriverPostgres || cfg.Addr != "localhost:5432" || cfg.Dbname != "ppanel" || cfg.Username != "postgres" {
		t.Fatalf("unexpected config: %+v", cfg)
	}

	dsn := Mysql{Config: *cfg}.Dsn()
	parsed, err := url.Parse(dsn)
	if err != nil {
		t.Fatalf("parse generated postgres dsn: %v", err)
	}
	if got := parsed.Query().Get("TimeZone"); got != "Asia/Shanghai" {
		t.Fatalf("postgres TimeZone = %q, want Asia/Shanghai", got)
	}
	if got := parsed.Query().Get("application_name"); got != defaultPostgresApplicationName {
		t.Fatalf("postgres application_name = %q, want %q", got, defaultPostgresApplicationName)
	}
	if cfg.ConnMaxLifetime != DefaultConnMaxLifetimeSeconds || cfg.ConnMaxIdleTime != DefaultConnMaxIdleTimeSeconds {
		t.Fatalf("postgres pool lifetimes = (%d, %d), want (%d, %d)",
			cfg.ConnMaxLifetime, cfg.ConnMaxIdleTime, DefaultConnMaxLifetimeSeconds, DefaultConnMaxIdleTimeSeconds)
	}
}

func TestPostgresDSNPreservesApplicationNameOverride(t *testing.T) {
	cfg := Config{
		Driver:   DriverPostgres,
		Addr:     "localhost:5432",
		Dbname:   "ppanel",
		Username: "postgres",
		Config:   "sslmode=disable&application_name=worker",
	}
	parsed, err := url.Parse((Mysql{Config: cfg}).Dsn())
	if err != nil {
		t.Fatalf("parse generated postgres dsn: %v", err)
	}
	if got := parsed.Query().Get("application_name"); got != "worker" {
		t.Fatalf("postgres application_name = %q, want worker", got)
	}
}

func TestPingMySQL(t *testing.T) {
	dsn := os.Getenv("PPANEL_TEST_MYSQL_DSN")
	if dsn == "" {
		t.Skip("set PPANEL_TEST_MYSQL_DSN to run MySQL/MariaDB ping test")
	}
	if !PingDatabase(DriverMySQL, dsn) {
		t.Fatal("mysql ping failed")
	}
}

func TestPingPostgres(t *testing.T) {
	dsn := os.Getenv("PPANEL_TEST_POSTGRES_DSN")
	if dsn == "" {
		t.Skip("set PPANEL_TEST_POSTGRES_DSN to run PostgreSQL ping test")
	}
	if !PingDatabase(DriverPostgres, dsn) {
		t.Fatal("postgres ping failed")
	}
}
