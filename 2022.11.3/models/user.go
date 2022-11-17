package models

import "database/sql/driver"
/*
1）占位符、tag、数据库字段，三者要一致，否则会冲突！！！
*/

type User struct { // 字段顺序怎么安排？类型和mysql怎么匹配？
    ID int
    Name string //`db:"name"`
    Age int //`db:"age"`
}
// Value 把属性用空接口切片的形式返回
func (u User) Value() (driver.Value, error){
    return []interface{}{u.Name,u.Age}, nil
}
