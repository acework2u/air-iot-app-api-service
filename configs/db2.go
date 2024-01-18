package configs

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func SmartConnect() (*bun.DB, error) {
	sqldb, err := sql.Open("mysql", "root:tiger@tcp(host.docker.internal:3306)/saijode_eapp")

	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectToMariaDB() (*sql.DB, error) {
	//db, err := sql.Open("mysql", "root:tiger@tcp(127.0.0.1:3306)/saijode_eapp")
	db, err := sql.Open("mysql", "root:tiger@tcp(host.docker.internal:3306)/saijode_eapp")
	//defer db.Close()
	if err != nil {
		return nil, err
	}
	// Ping the MariaDB server to ensure connectivity
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
