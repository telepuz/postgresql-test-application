package postgresql

import (
	"context"
	"fmt"
	"time"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	READ_QUERY  = "SELECT t FROM test_table ORDER BY 1 DESC LIMIT 1;"
	WRITE_QUERY = "INSERT INTO test_table(t) VALUES($1);"
)

type Postgresql struct {
	Name string
	conn *pgxpool.Pool
}

func NewPostgresql(name string, c *pgxpool.Pool) *Postgresql {
	return &Postgresql{
		Name: name,
		conn: c,
	}
}

func (p *Postgresql) Read() (string, error) {
	var res string
	cx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := p.conn.QueryRow(cx, READ_QUERY).Scan(&res)
	if err != nil {
		return "", fmt.Errorf("ReadQuery failed: %v", err)
	}
	return res, err
}

func (p *Postgresql) Write() (string, error) {
	val := time.Now().Format("2006-01-02T15:04:05.000")
	cx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := p.conn.QueryRow(
		cx,
		WRITE_QUERY, val).Scan()
	if err != pgx.ErrNoRows {
		return "", fmt.Errorf("trying write: %s, readquery failed: %s", val, err.Error())
	}
	return val, nil
}
