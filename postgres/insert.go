package postgres

import (
	"context"
	"fmt"
)

func (c *ConnectData) Insert(jsonb []byte, service string) error {
	_, err := c.conn.Exec(context.Background(), `INSERT INTO configs (conf) VALUES
        ($1)`, jsonb)
	if err != nil {
		return fmt.Errorf("DB ERROR : %v", err)
	}
	return nil
}
