package sqlx

import (
    "fmt"
    "hpytest/sqldemo1/models"
    "log"
)

// QueryAUserByID 查询一个用户; 用 get 代替 QueryRow
func QueryAUserByID(id int) (user models.User, err error) {
    sqlStr := "SELECT name,age FROM user WHERE id=?"
    //row := db.QueryRow(sqlStr, id)
    //err = row.Scan(&user.Name, &user.Age)
    //if err != nil {
    //    log.Println(err)
    //    return user, err
    //}
    //return // user 不是nil，上级必须先判断 err，才知道 user 是否可用
    db.Get(&user,sqlStr,2)
}

// QueryUsersByID 查询多个用户
func QueryUsersByID() {

}









/*
NamedExec   NamedQuery
可以先指定占位符的名字（对应 map 的字段名），然后在执行时通过 map 传参
map[string]interface{}   value 是 interface{} 类型
sql 占位符的单词要和 map 的 key 一样
如果不用map用struct，那么就要用 tag  `db:"name"` 这种形式，让占位符知道自己对应谁
*/
func insertUserDemo()(err error){
    sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
    userMap := map[string]interface{}{
        "name": "七米",
        "age": 28,
    }
    _, err = db.NamedExec(sqlStr, userMap)
    return
}

func namedQuery(){
    sqlStr := "SELECT * FROM user WHERE name=:name"
    // 使用map做命名查询
    rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "七米"})
    if err != nil {
        fmt.Printf("db.NamedQuery failed, err:%v\n", err)
        return
    }
    defer rows.Close()
    for rows.Next(){
        var u models.User
        err := rows.StructScan(&u)
        if err != nil {
            fmt.Printf("scan failed, err:%v\n", err)
            continue
        }
        fmt.Printf("user:%#v\n", u)
    }

    u := models.User{
        Name: "七米",
    }
    // 使用结构体命名查询，根据结构体字段的 db tag进行映射
    rows, err = db.NamedQuery(sqlStr, u)
    if err != nil {
        fmt.Printf("db.NamedQuery failed, err:%v\n", err)
        return
    }
    defer rows.Close()
    for rows.Next(){
        var u models.User
        err := rows.StructScan(&u)
        if err != nil {
            fmt.Printf("scan failed, err:%v\n", err)
            continue
        }
        fmt.Printf("user:%#v\n", u)
    }
}
