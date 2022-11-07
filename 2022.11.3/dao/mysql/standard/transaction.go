package standard

import (
    "fmt"
    "log"
)

//Begin Commit Rollback

func TransactionDemo() {
    tx,err := db.Begin()
    if err != nil {
        log.Println(err)
        return
    }
    sqlStr := "UPDATE user SET age=30 WHERE id=?"
    res, err := tx.Exec(sqlStr, 3)
    if err != nil {
        tx.Rollback()
        log.Println(err)
        return
    }
    n, err := res.RowsAffected()
    if err != nil {
        tx.Rollback()
        log.Println(err)
        return
    }
    fmt.Println(n)
    if n == 1 {
        tx.Commit()
        log.Println("commit success")
    } else {
        tx.Rollback()
        log.Println("xxx fail, rollback")
    }
}
