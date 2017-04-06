package data

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	// Database driver is not used directly
	_ "github.com/go-sql-driver/mysql"
)

type singleton struct {
	Connection *sql.DB
}

var instance *singleton
var once sync.Once

// GetInstance return DB instance (singleton)
func GetInstance() *singleton {
	once.Do(func() {
		config := GetConfigInstance()

		dsn := fmt.Sprintf("%v:%v@/%v", config.Configuration.Db.Username, config.Configuration.Db.Password, config.Configuration.Db.Database)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		}
		instance = &singleton{Connection: db}
	})
	return instance
}
