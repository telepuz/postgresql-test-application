package app

import (
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/telepuz/postgresql-test-application/internal/postgresql"
)

type AppContext struct {
	PostgresqlRW *postgresql.Postgresql
	PostgresqlRO *postgresql.Postgresql
}

func Run(c *AppContext) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for {
			res, err := c.PostgresqlRW.Write()
			if err != nil {
				slog.Error(err.Error())
			}
			slog.Info(fmt.Sprintf("Write: %s", res))
			time.Sleep(500 * time.Millisecond)
		}
		wg.Done()
	}()

	go func() {
		for {
			res, err := c.PostgresqlRO.Read()
			if err != nil {
				slog.Error(err.Error())
			}
			slog.Info(fmt.Sprintf("Read: %s", res))
			time.Sleep(500 * time.Millisecond)
		}
		wg.Done()
	}()

	wg.Wait()
	slog.Info("Complete all go routines. Exit...")
}
