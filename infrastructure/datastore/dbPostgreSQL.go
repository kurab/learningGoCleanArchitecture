package datastore

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "api/config"
)

func NewPostgreSQL() *gorm.DB {
    connectString := fmt.Sprintf(
        "host=%s port=%s user=%s dbname= %s password=%s %s",
        config.C.Postgres.HOST,
        config.C.Postgres.PORT,
        config.C.Postgres.USER,
        config.C.Postgres.DBNAME,
        config.C.Postgres.PASS,
        config.C.Postgres.OPTION,
    )
    db, err := gorm.Open(config.C.Postgres.DBMS, connectString)

    if err != nil {
        panic(err.Error())
    }

    db.LogMode(true)

    return db
}
