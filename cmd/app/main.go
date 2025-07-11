package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/telepuz/postgresql-test-application/internal/app"
	"github.com/telepuz/postgresql-test-application/internal/logger"
	"github.com/telepuz/postgresql-test-application/internal/postgresql"
)

func main() {
	envFile, _ := godotenv.Read(".env")

	err := logger.ConfigureSlog()
	if err != nil {
		slog.Error(
			fmt.Sprintf("main(): %s", err))
		os.Exit(1)
	}

	rw_conn, err := pgxpool.Connect(
		context.Background(),
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			envFile["DB_USER"],
			envFile["DB_PASS"],
			envFile["DB_HOST"],
			envFile["DB_PORT_RW"],
			envFile["DB_NAME"],
		))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	ro_conn, err := pgxpool.Connect(
		context.Background(),
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			envFile["DB_USER"],
			envFile["DB_PASS"],
			envFile["DB_HOST"],
			envFile["DB_PORT_RO"],
			envFile["DB_NAME"],
		))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	pg_rw := postgresql.NewPostgresql(
		"test-cluster",
		rw_conn)
	pg_ro := postgresql.NewPostgresql(
		"test-cluster",
		ro_conn)
	c := app.AppContext{
		PostgresqlRW: pg_rw,
		PostgresqlRO: pg_ro,
	}

	app.Run(&c)
}
