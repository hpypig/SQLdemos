package main

import (
    "hpytest/sqldemo1/dao/mysql" // 这个和mysql驱动重名了，但不在一个包内不影响
    "log"
)

// 要学会文件内（项目内）全局改变量名

func main() {
    err := mysql.InitDB()
    if err != nil {
        log.Println(err.Error()) // Println 可以直接传 err，有什么区别，之后自己实现一个 error 试一下
    }
    // 连接成功才注册关闭函数
    defer mysql.Close()
}
