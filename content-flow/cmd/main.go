package main

import (
	"content-flow/internal/process"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
)

func main() {
	mysqlDB, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)

	process.ExecContentFlow(mysqlDB)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("Shutting down server...")
}
