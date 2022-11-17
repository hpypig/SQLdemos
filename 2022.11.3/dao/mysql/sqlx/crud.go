package sqlx

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    "hpytest/sqldemo1/models"
    "log"
    "strings"
)
/*
新增api
Get 可以直接存到结构体里
NamedQuery 也是返回多行 rows 但可以用struct、map传占位符
    map[string]interface{}{"name":name}
    placeholder := struct {
      Name string // 这里的属性首字母必须大写，否则下面 rows 会处 nil 错误；原理待了解
    }{
      Name: name,
    }
NamedExec 用 struct 传占位符


 */

// 查询 -----------------------------------------------------------------------

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
    err = db.Get(&user, sqlStr,2)
    if err != nil {
        log.Println(err)
        return
    }
    return
}

// QueryUsersByID 查询多个用户
func QueryUsersByID(id int) (users []models.User, err error) {
    sqlStr := "SELECT id, name, age FROM USER WHERE id>?" // 如果字段不全会发生什么？？？？？？
    err = db.Select(&users, sqlStr, id)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(users)
    return
}

func NamedQueryUserSByName(name string) (err error){
    sqlStr := "SELECT * FROM user WHERE name=:name" // 用 name 占位，表示需要的参数
    // 下面就应该知道 name 是多少 ——> 用 map 传参
    // 这个是匿名map？
    rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name":name})
    if err != nil {
        log.Println(err)
        return
    }
    defer rows.Close() // rows 要保证关闭，why？？？
    for rows.Next() {
        var u models.User
        // 可以直接传结构体了
        // 而且是传递指针进去，而不是返回一个结构体
        // 别人是怎么传值可以好好研究
        err = rows.StructScan(&u)
        if err != nil {
            log.Println(err)
            return
        }
        fmt.Println(u)
    }
    // --------------
    // 也可以用 struct 传参
    placeholder := struct {
      Name string // 这里的属性首字母必须大写，否则下面 rows 会处 nil 错误；原理待了解
    }{
      Name: name,
    }

    rows, err = db.NamedQuery(sqlStr, placeholder)
    defer rows.Close()
    for rows.Next() {
       var u models.User
       // 可以直接传结构体了
       // 而且是传递指针进去，而不是返回一个结构体
       // 别人是怎么传值可以好好研究
       err = rows.StructScan(&u)
       if err != nil {
           log.Println(err)
           return
       }
       fmt.Println(u)
    }
    return
}

// 删改查 -----------------------------------------------------------------------

// InsertUserByNamedExec 的
func InsertUserByNamedExec(user models.User) (err error) {
    // 默认用 struct 字段名的小写匹配占位符，除非用 tag 指明
    sqlStr := "INSERT INTO user(name, age) VALUES (:name, :age)"
    res, err := db.NamedExec(sqlStr,&user) // `db:"name"`
    if err != nil {
        log.Println(err)
        return
    }
    id, err := res.LastInsertId()
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(id)
    return
}
// sqlx.In 批量插入 ----------------------------------------------
// sql 里的 (?,?,?) 可以先由 (?) 代替，执行后被替换成正确的sql语句
// 注：insert 最后连续跟多个(?,?) 是框架特有的，不是sql的语法

// InsertUsers 必须先实现传入的结构体的 Value 接口；假如有多种插入字段组合，Value咋搞？
func InsertUsers(users []interface{}) (err error) { // 要用 interface{} 切片而不是 []User
    sqlStr := "INSERT INTO user (name, age) VALUES(?), (?), (?)"
    query, args, err := sqlx.In(sqlStr,users...)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(query) // INSERT INTO user (name, age) VALUES(?, ?), (?, ?), (?, ?)
    fmt.Println(args) // [rr 1 tt 2 qq 3]
    // 相当于直接用结构体字段，帮我们进行了 args 生成
    //（直接用结构体而无需自己提取字段后 append args）
    res, err := db.Exec(query, args...)
    if err != nil {
        log.Println(err)
        return
    }
    id, err := res.LastInsertId()
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(id)
    n, err := res.RowsAffected()
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(n)
    return
}

// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []*models.User) error { // 感觉这个简洁点啊，不用实现接口，也不用传空接口切片
    _, err := db.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
    return err
}

// QueryUsersByIDs sql in 查询
func QueryUsersByIDs(ids []int) (users []models.User, err error){
    sqlStr := "SELECT name, age FROM user WHERE id IN (?)"
    query, args, err := sqlx.In(sqlStr,ids)
    if err != nil {
        log.Println(err)
        return
    }
    //fmt.Println(query) // SELECT name, age FROM user WHERE id IN (?, ?, ?, ?, ?)
    //fmt.Println(args) // [1 2 3 4 5]
    // sql 老方法 Query ——> rows ——> u
    //rows, err := db.Query(query, args...)
    //defer rows.Close()
    //for rows.Next() {
    //    var u models.User
    //    err = rows.Scan(&u.Name, &u.Age)
    //    if err != nil {
    //        log.Println(err)
    //        return
    //    }
    //    fmt.Println(u)
    //}
    //query = db.Rebind(query) // 好像不绑也行，不知道为啥
    //fmt.Println(query) // SELECT name, age FROM user WHERE id IN (?, ?, ?, ?, ?)
    err = db.Select(&users,query,args...) // `db:"Name"` 会报错，列 name 在 User 找不到对应属性
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(users)
    return
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int)(users []models.User, err error){
    // 动态填充id
    strIDs := make([]string, 0, len(ids))
    for _, id := range ids {
        strIDs = append(strIDs, fmt.Sprintf("%d", id))
    }

    query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
    if err != nil {
        log.Println(err)
        return
    }
    //fmt.Println(query) //SELECT name, age FROM user WHERE id IN (?, ?, ?, ?, ?) ORDER BY FIND_IN_SET(id, ?)
    fmt.Println(args)
    // sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
    //query = db.Rebind(query)

    err = db.Select(&users, query, args...)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(users)
    return
}






















//------------------------------
//-------------------------下面是七米写的
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
