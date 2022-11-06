package postgres

import (
	"context"
	"fmt"
)

func (c *ConnectData) Insert(jsonb []byte, service string) error {
	_, err := c.conn.Exec(context.Background(), `INSERT INTO configs (conf, services) VALUES
        ($1, $2)`, jsonb, service)
	if err != nil {
		return fmt.Errorf("DB ERROR : %v", err)
	}
	return nil
}
