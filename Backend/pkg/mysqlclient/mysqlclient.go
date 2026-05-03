package mysqlclient

import (
	"bibliotheca/internal/config"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySqlClient(cfg *config.Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to Database : %v", err)
	}
	
	db.SetMaxOpenConns(25)              // max open connections to DB
	db.SetMaxIdleConns(10)              // max idle connections kept in pool
	db.SetConnMaxLifetime(5 * time.Minute)  // recycle connections every 5 min
	db.SetConnMaxIdleTime(2 * time.Minute)  // close idle connections after 2 min


	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error pinging database: %v", err)
	}

	log.Println("Database connected successfully !!!")
	return db, nil
}