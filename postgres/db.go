package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type ConnectData struct {
	conn *pgx.Conn
}

func NewConnStruct(url string) *ConnectData {
	conn, err := Connect(url)
	if err != nil {
		fmt.Println("failed connect db")
		return nil
	}
	return &ConnectData{conn: conn}
}

func Connect(url string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("failed get pgx conn : %v", err)
	}
	return conn, nil
}

func (c *ConnectData) InitDbTables() error {
	_, err := c.conn.Exec(context.Background(), `CREATE TABLE configs (
    id serial PRIMARY KEY,
    conf JSONB  NOT NULL
	);
`)
	if err != nil {
		if err.Error() != `ERROR`+": "+`relation "configs" already exists`+" (SQLSTATE "+`42P07`+")" {
			return fmt.Errorf("Failed Create table configs : %v", err)
		}
	}

	_, err = c.conn.Exec(context.Background(), `CREATE TABLE services (
    id serial PRIMARY KEY,
   	service VARCHAR(40) NOT NULL,
	config JSONB
	);
`)
	if err != nil {
		if err.Error() != `ERROR`+": "+`relation "services" already exists`+" (SQLSTATE "+`42P07`+")" {
			return fmt.Errorf("Failed Create table configs : %v", err)
		}
	}

	return nil
}
