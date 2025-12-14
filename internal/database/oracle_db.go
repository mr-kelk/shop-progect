package database

import (
	"database/sql"
	"example/shop-progect/config"

	go_ora "github.com/sijms/go-ora/v2"
)

var DB *sql.DB

func Init() error {
	connStr := go_ora.BuildUrl(
		config.Cfg.OracleHost,
		config.Cfg.OraclePort,
		config.Cfg.OracleService,
		config.Cfg.OracleUser,
		config.Cfg.OraclePass,
		nil,
	)

	var err error
	DB, err = sql.Open("oracle", connStr)
	if err != nil {
		return err
	}

	return DB.Ping()
}
