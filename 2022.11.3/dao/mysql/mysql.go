package mysql

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql" // goland 有个sync的过程，Linux上呢？tidy？get？
    "hpytest/sqldemo1/models"
    "log"
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

func QueryUser(id int, user *models.User) (err error){
    sqlStr := "select id, name, age from `user` where id=?"
    err = db.QueryRow(sqlStr, id).Scan(&user.ID, &user.Name, &user.Age)
    if err != nil {
        log.Println("QueryUser fail") // 我要怎么记录：是哪个文件，哪行代码出错？
        return
    }
    return nil
}
func QueryUserMoreThanID(id int) (err error){
    sqlStr := "select * from user where id>?"
    rows, err := db.Query(sqlStr, id)
    if err != nil {
        log.Println(err)
    }
    // 释放数据库连接
    defer rows.Close()
    for rows.Next() {
        var u models.User
        err = rows.Scan(&u.ID, &u.Name, &u.Age)
        if err != nil {
            log.Println(err)
            return
        }
        fmt.Println(u)
    }
    return
}

func InsertUser(user models.User) (err error) { // 是不是只要可能出现error，我就要返回 error？
    sqlStr := "INSERT INTO `user`(name, age) values(?,?)"
    ret, err := db.Exec(sqlStr, user.Name, user.Age)
    if err != nil {
        log.Println("InsertUser fail1")
        return
    }
    id, err := ret.LastInsertId() // 有自增列才有
    if err != nil { // 这里的err是因为什么？
        log.Println("InsertUser fail2")
        return
    }
    fmt.Println(id)
    return
}

// 这个函数没有定参数，不知道怎么定。要用到设计模式吗？为了适应update不同字段的情况

func UpdateUserXXXByXXX() (err error){ // 每种修改都要写个接口吗？如果用 != 默认值判断，那如果就是要改成默认值呢
    sqlStr := "UPDATE user SET age=? where id=?"
    age := 3
    id := 0
    ret, err := db.Exec(sqlStr, age, id)
    if err != nil {
        log.Println(err)
    }
    n, err := ret.RowsAffected()
    if err != nil {
        log.Println(err)
    }
    fmt.Println(n)
    return // 如果最开始已经打印过 err 了，上一层调用还要处理吗
}

func DeleteUserByID(id int) (err error){
    sqlStr := "DELETE FROM `user` WHERE ID=?"
    ret, err := db.Exec(sqlStr, id)
    if err != nil {
        log.Println(err)
    }
    n, err := ret.RowsAffected()
    if err != nil {
        log.Println(err)
    }
    fmt.Println(n)
    return
}


