package mysql

import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql" // goland 有个sync的过程，Linux上呢？tidy？get？
)

var db *sql.DB
func InitDB() (err error){
    address := "root:hh424@tcp(127.0.0.1:3306)/sql_demos"
    db, err = sql.Open("mysql",address)
    if err != nil {
        // panic(err)
        return
    }
    err = db.Ping()
    return
}
func Close() {
    db.Close()
}
