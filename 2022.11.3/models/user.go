package models
type User struct { // 字段顺序怎么安排？类型和mysql怎么匹配？
    ID int
    Name string `db:"Name"`
    Age int
}
