package configs

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Db3Connect() (*gorm.DB, error) {
	var err error
	dsn := "root:tiger@tcp(host.docker.internal:3306)/saijode_smartapp?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("No connect DB")
		return nil, err
	}

	return db, nil
}
