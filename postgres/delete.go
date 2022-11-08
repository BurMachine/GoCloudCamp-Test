package postgres

import (
	"context"
	"fmt"
)

func (c *ConnectData) Delete(conf []byte) error {

	pgx, err := c.conn.Exec(context.Background(), `SELECT * FROM services WHERE config=$1`, conf)
	if err != nil {
		return fmt.Errorf("DB ERROR : %v", err)
	}
	if pgx.String() == "SELECT 0" {
		_, err = c.conn.Exec(context.Background(), `DELETE FROM configs WHERE conf=$1`, conf)
		if err != nil {
			return fmt.Errorf("DB REMOVE ERROR : %v", err)
		}
		var tmp []byte
		_, err := c.conn.Exec(context.Background(), `UPDATE services SET config=$1 WHERE config=$2`, tmp, conf)
		if err != nil {
			return fmt.Errorf("DB ERROR : %v", err)
		}
	} else {
		return fmt.Errorf("DB ERROR : config in use")
	}
	return nil
}
