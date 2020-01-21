package datastore

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMySQL() *gorm.DB {
    connectString := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s%s",
        "user",
        "secret",
        "0.0.0.0",
        "3306",
        "goSample",
        "?charset=utf8&parseTime=True&loc=Local",
    )
    db, err := gorm.Open("mysql", connectString)
    if err != nil {
        panic(err.Error())
    }
    db.LogMode(true)
    db.Set("gorm:table_options", "ENGIN=InnoDB")

    return db
}
