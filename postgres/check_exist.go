package postgres

import (
	"context"
)

func (c *ConnectData) CheckConf(jsonb []byte) bool {
	pgx, err := c.conn.Exec(context.Background(), `SELECT * FROM configs WHERE conf=$1`, jsonb)
	if err != nil {
		return true
	}
	if pgx.String() == "SELECT 0" {
		return false
	}
	return true
}
