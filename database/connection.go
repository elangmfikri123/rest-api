package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDbInstance() *gorm.DB {
	// dsn := "host=localhost user=postgres password=12345 dbname=penjualan port=5432 sslmode=disable"
	dsn := "host=localhost user=root password=nopalgemay32.! dbname=penjualan port=5432 sslmode=disable"

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Gagal Terhubung Database")
	}

	return db
}
