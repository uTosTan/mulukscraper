package data

import (
    "sync"
    "encoding/json"
    "os"
    "log"
)

type Configuration struct {
    Db Database
}

type Database struct {
    Host        string
    Database    string
    Username    string
    Password    string
}

type configSingleton struct {
    Configuration *Configuration
}

var configInstance *configSingleton
var one sync.Once

func GetConfigInstance() *configSingleton {
    one.Do(func() {
        file, _ := os.Open("config.json")
        decoder := json.NewDecoder(file)
        configuration := Configuration{}
        err := decoder.Decode(&configuration)

        if err != nil {
            log.Fatal(err)
        }

        configInstance = &configSingleton{ Configuration: &configuration }
    })
    return configInstance
}