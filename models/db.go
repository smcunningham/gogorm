package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() *gorm.DB {
	var db *gorm.DB

	var err error
	dataSourceName := "root:D20!@shboard@tcp(localhost:3306)/gogormdb?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	// Create the database. This is a one-time step
	// Comment out if running multiple tines- You may see an error otherwise
	//db.Exec("CREATE DATABASE orders_db")
	//db.Exec("USE orders_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&Order{}, &Item{})

	return db
}
