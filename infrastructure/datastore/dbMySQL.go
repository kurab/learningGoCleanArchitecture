package datastore

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "api/config"
)

func NewMySQL() *gorm.DB {
    connectString := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s%s",
        config.C.Mysql.USER,
        config.C.Mysql.PASS,
        config.C.Mysql.HOST,
        config.C.Mysql.PORT,
        config.C.Mysql.DBNAME,
        config.C.Mysql.OPTION,
    )
    db, err := gorm.Open(config.C.Mysql.DBMS, connectString)
    if err != nil {
        panic(err.Error())
    }
    db.LogMode(true)
    db.Set("gorm:table_options", "ENGIN=InnoDB")

    return db
}
