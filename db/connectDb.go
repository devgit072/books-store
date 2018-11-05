package db

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

const (
	USER_NAME="<>"
	PASSWORD="<>"
	host="10.2.41.142"
)

func ConnectDB() (*sql.DB, error) {

	connStr := fmt.Sprintf("user=%s host=%s dbname=books_db sslmode=disable password=%s port=5432",
		host,USER_NAME, PASSWORD)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		glog.Fatal(err)
		return nil, err
	}
	//err = db.Ping()

	return db, nil
}
