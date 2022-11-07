package sqlx

import (
    "errors"
    //"fmt"
    "log"
)

func TransactionDemo() (err error){ // 通过 err 知道是哪步错了
    tx, err := db.Begin()
    if err != nil {
        log.Println(err)
        return
    }
    // panic 和 err 下都要 Rollback，但可以最后执行
    // defer 最佳实践是什么？这样会影响性能吗？
    defer func() {
        if p := recover(); p!=nil {
            tx.Rollback()
            panic(p) // 先 rollback 再重新panic
        } else if err != nil {
            log.Println("rollback")
            tx.Rollback()
        } else {
            err = tx.Commit() // 最后这个 err 由上层调用处理
            log.Println("commit")
        }
    }()

    sqlStr1 := "UPDATE user SET age=50 WHERE id>?"
    res, err := tx.Exec(sqlStr1, 1)
    if err != nil {
        log.Println(err)
        return err
    }
    n, err := res.RowsAffected()
    if err != nil {
        log.Println(err)
        return err
    }
    if n==0 {
        return errors.New("n==0,sql1 failed")
    }
    sqlStr2 := "UPDATE user set age=60 where i=?"
    res, err = tx.Exec(sqlStr2, 2)
    if err != nil {
        log.Println(err)
        return err
    }
    n, err = res.RowsAffected()
    if err != nil {
        log.Println(err)
        return err
    }
    if n!=1 {
        return errors.New("n!=1,sql1 failed")
    }
    return err
}
