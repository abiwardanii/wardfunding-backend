package main

import (
	"log"
	"net/http"
	"wardfunding/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/wardstartup?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	  if err != nil {
		log.Fatal(err.Error())
	  }
	  userRepository := user.NewRepository(db)
	  userService := user.NewService(userRepository)
}