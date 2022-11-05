package main

import (
    "fmt"
    "hpytest/sqldemo1/dao/mysql" // 这个和mysql驱动重名了，但不在一个包内不影响
    "hpytest/sqldemo1/models"
    "log"
)
// 参考：https://www.liwenzhou.com/posts/Go/go_mysql/
// 要学会文件内（项目内）全局改变量名

func main() {

    testMySQL()

}

func testMySQL() {
    err := mysql.InitDB()
    if err != nil {
        log.Println(err.Error()) // Println 可以直接传 err，有什么区别，之后自己实现一个 error 试一下
    }
    // 连接成功才注册关闭函数
    defer mysql.Close()
    //testInsert()
    //testQueryRow() // QueryRow ——> rows ——> rows.Next() rows.Scan()
    //testQueryRows() // Query ——> row ——> row.Scan()
    //testDelete()   // Exec
    // update 没确定咋写，反正也是 Exec
}


func testInsert() (err error){
    // 插入
    user1 := models.User{ // 插入时直接实例（防逃逸？）
        Name:"aaa",
        Age:1,
    }
    user2 := models.User{
        Name:"aaa",
        Age:2,
    }
    err = mysql.InsertUser(user1)
    if err!=nil {
        log.Println(err)
    }
    err = mysql.InsertUser(user2)
    if err != nil {
        log.Println(err)
    }
    return
}
func testQueryRow() (err error) {
    // 查询一行结果
    id := 1
    userA := new(models.User) // 获取结果时用指针？
    err = mysql.QueryUser(id, userA)
    if err != nil {
        log.Println(err)
    }
    fmt.Println(userA)
    return
}
func testQueryRows() {
    mysql.QueryUserMoreThanID(0)
}


func testUpdate() (err error) {
    return
}
func testDelete() {
    mysql.DeleteUserByID(1)
}
