package banco

import (
	"fmt"
	"log"
	"sisvac/utils"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func GetConnection() (*sqlx.DB, error) {
	// Load vars env
	dbUser := "root"
	dbPwd := ""
	dbName := "base_de_dados"
	dbAddress := "localhost"
	dbPort := "3306"

	// init database
	dataSource := fmt.Sprintf("mysql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPwd, dbAddress, dbPort, dbName)
	db, err := sqlx.Open("mysql", dataSource)

	if err != nil {
		log.Print("Error while accessing database: " + err.Error())
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

conx, err := GetConnection()

if err != nil {
	fmt.Println("ERRO!", err)
}
