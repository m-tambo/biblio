package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

var (
	db *sql.DB
	host		= viperConfigVariable("DB_HOST")
	port, _ 	= strconv.Atoi(viperConfigVariable("DB_PORT"))
	user     	= viperConfigVariable("DB_USER")
	password 	= viperConfigVariable("DB_PASS")
	dbname   	= viperConfigVariable("DB_NAME")
)

func NewDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func viperConfigVariable(key string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".") // look for config in the working directory

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}