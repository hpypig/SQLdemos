package TestMySQL

import (
    "hpytest/sqldemo1/dao/mysql/sqlx"
    "log"
)

func SqlxMain() {
    err := sqlx.InitDB()
    if err != nil {
        log.Println(err)
        return
    }
    defer func() {
        err = sqlx.Close() // 什么时候不能用闭包变量？？？？
        if err != nil {
            log.Println(err)
        }
    }()
    //-------------------------------
    // 为什么李要用 new 先创建 user，再获取呢
    //user,err := sqlx.QueryAUserByID(2)
    //if err != nil {
    //    log.Println(err)
    //    return // err 是不是要一直返回啊？
    //}
    //fmt.Println(user) // {0 aaa 10} 因为没有查 id 所以是默认值，要注意默认值问题！！
    //---------------------
    //sqlx.TransactionDemo()
    //sqlx.NamedQueryUserSByName("aaa")
    //sqlx.QueryUsersByID(1)

    //user := models.User{Name: "ww", Age: 41}
    //sqlx.InsertUserByNamedExec(user)

    //user1 := models.User{Name: "rr", Age:1 }
    //user2 := models.User{Name: "tt", Age: 2}
    //user3 := models.User{Name: "qq", Age: 3}
    //var users []interface{}
    //users = append(users, user1, user2, user3)
    //sqlx.InsertUsers(users)

    ids := []int{3,2,4,5,6}
    //sqlx.QueryUsersByIDs(ids) // [{0 aaa 10} {0 aaa 30} {0 aaa 2} {0 zs 21}]
    sqlx.QueryAndOrderByIDs(ids) //[{0 aaa 10} {0 aaa 30} {0 aaa 2} {0 zs 21}]
}

