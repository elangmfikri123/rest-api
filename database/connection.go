package database

import (
	"fmt"
	"log"

	"github.com/elangmfikri123/rest-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDbInstance() *gorm.DB {
	dbConfig := config.DbConfig()

	// dsn := "host=localhost user=postgres password=12345 dbname=penjualan port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", //sslmode=%s
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DbName,
		dbConfig.Port,
		//dbConfig.SslMode,
	)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Gagal Terhubung Database")
	}

	return db
}
