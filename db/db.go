package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/saskaradit/echo-test/util"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.SSLMode)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully Connected!")
}
