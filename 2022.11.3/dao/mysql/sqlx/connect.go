package sqlx
//ref: https://www.liwenzhou.com/posts/Go/sqlx/
import (
	"fmt"
    "github.com/jmoiron/sqlx"
    "log"
)
// sqlx 兼容了 sql，即 sql 原来的方法也可以用
var db *sqlx.DB

func InitDB() (err error) {
    dsn := "root:hh424@tcp(127.0.0.1:3306)/sql_demos?charset=utf8mb4&parseTime=True"
    // 也可以使用MustConnect连接不成功就panic
    db, err = sqlx.Connect("mysql", dsn)
    if err != nil {
        fmt.Printf("connect DB failed, err:%v\n", err)
        return
    }
    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(10)
    return
}
func Close() (err error){
    err = db.Close()
    if err != nil {
        log.Println(err)
        //return
    }
    return
}
