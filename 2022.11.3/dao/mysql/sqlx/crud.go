package sqlx

import (
    "fmt"
    "hpytest/sqldemo1/models"
)

//ref: https://www.liwenzhou.com/posts/Go/sqlx/

// NamedExec   NamedQuery
// 可以先指定占位符的名字（对应 map 的字段名），然后在执行时通过 map 传参
// map[string]interface{}   value 是 interface{} 类型
// sql 占位符的单词要和 map 的 key 一样
// 如果不用map用struct，那么就要用 tag  `db:"name"` 这种形式，让占位符知道自己对应谁

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
