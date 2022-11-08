package postgres

import (
	"context"
	"fmt"
)

func (c *ConnectData) GetConfFromService(service string) ([]byte, error) {
	var conf []byte
	err := c.conn.QueryRow(context.Background(), "select config from services where service=$1", service).Scan(&conf)
	if err != nil {
		return nil, fmt.Errorf("DB ERROR : %v", err)
	}
	return conf, nil
}
