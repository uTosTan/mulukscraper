package data

import (
    "sync"
    "database/sql"
    _"github.com/go-sql-driver/mysql"
    "log"
)

type singleton struct {
    Connection *sql.DB
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        db, err := sql.Open("mysql", "scrapenepal:suraj!2@/scrapenepal")
        if err != nil {
            log.Fatal(err)
        }
        instance = &singleton{ Connection: db }
    })
    return instance
}