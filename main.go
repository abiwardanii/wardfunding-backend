package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/wardstartup?charset=utf8mb4&parseTime=True&loc=Local"
  	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	  if err != nil {
		panic("failed to connect database")
	  }

	fmt.Println("Database Connected")
}