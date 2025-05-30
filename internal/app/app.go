package app

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/telepuz/postgresql-test-application/internal/postgresql"
)

type AppContext struct {
	Postgresql *postgresql.Postgresql
}

func Run(c *AppContext) {
	for {
		res, err := c.Postgresql.Write()
		if err != nil {
			slog.Error(err.Error())
		}
		slog.Info(fmt.Sprintf("Write: %s", res))

		res, err = c.Postgresql.Read()
		if err != nil {
			slog.Error(err.Error())
		}
		slog.Info(fmt.Sprintf("Read: %s", res))

		time.Sleep(500 * time.Millisecond)
	}
}
