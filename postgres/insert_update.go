package postgres

import (
	"context"
	"fmt"
)

func (c *ConnectData) InsertConf(jsonb []byte, service string) error {
	_, err := c.conn.Exec(context.Background(), `INSERT INTO configs (conf, serviceName) VALUES
        ($1, $2)`, jsonb, service)
	if err != nil {
		return fmt.Errorf("DB ERROR : %v", err)
	}
	return nil
}

func (c *ConnectData) UpdateService(jsonb []byte, service string) error {
	_, err := c.conn.Exec(context.Background(), `UPDATE services SET config=$1 WHERE service=$2`, jsonb, service)
	if err != nil {
		return fmt.Errorf("DB ERROR : %v", err)
	}
	fmt.Printf("Service '%s' config updated!", service)
	return nil
}
