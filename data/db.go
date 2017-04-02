package data

import (
    "sync"
    "database/sql"
    "log"
    "fmt"
)

type singleton struct {
    Connection *sql.DB
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        config := GetConfigInstance()

        dsn := fmt.Sprintf("%v:%v@/%v", config.Configuration.Db.Username, config.Configuration.Db.Password, config.Configuration.Db.Database)

        db, err := sql.Open("mysql", dsn)
        if err != nil {
            log.Fatal(err)
        }
        instance = &singleton{ Connection: db }
    })
    return instance
}