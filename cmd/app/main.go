package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type User struct {
	Id       int
	Email    string
	Password string
}

func main() {
	fmt.Println("APP APIs")
	var err error
	db, err = sql.Open("mysql", "root:tiger@tcp(127.0.0.1:3306)/saijode_eapp")
	if err != nil {
		panic("database error")
	}

	user, err := GetUserById(1)
	if err != nil {
		log.Fatal(err)
		//panic("Data not fault")
	}

	fmt.Println(user)

}

func GetUserById(id int) (*User, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id,email,password from cus_mstr where id =?"

	row := db.QueryRow(query, id)

	user := User{}

	err = row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func AddUser(user User) error {

	query := "insert into cus_mstr(email,password) values (?,?)"
	db.Exec(query, user.Email, user.Password)

	return nil
}
