package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=tp2u6jQbdM dbname=deneme sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	dbase := db.DB()
	defer dbase.Close()

	err = dbase.Ping()

	if err != nil {
		panic(err.Error())
	}
	println("Database Bağlantısı Başarılı")
}
