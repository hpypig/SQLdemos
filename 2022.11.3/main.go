package main

import (
    //"fmt"
    //"hpytest/sqldemo1/dao/mysql" // 这个和mysql驱动重名了，但不在一个包内不影响
    "hpytest/sqldemo1/logic/TestMySQL"
    //"hpytest/sqldemo1/models"
    //"log"
)
// 参考：https://www.liwenzhou.com/posts/Go/go_mysql/
// 要学会文件内（项目内）全局改变量名

func main() {

    TestMySQL.Main()

}

