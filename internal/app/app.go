package app

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/telepuz/postgresql-test-application/internal/postgresql"
)

type AppContext struct {
	PostgresqlRW *postgresql.Postgresql
	PostgresqlRO *postgresql.Postgresql
}

func Run(c *AppContext) {
	for {
		res, err := c.PostgresqlRW.Write()
		if err != nil {
			slog.Error(err.Error())
		}
		slog.Info(fmt.Sprintf("Write: %s", res))

		res, err = c.PostgresqlRO.Read()
		if err != nil {
			slog.Error(err.Error())
		}
		slog.Info(fmt.Sprintf("Read: %s", res))

		time.Sleep(500 * time.Millisecond)
	}
}
