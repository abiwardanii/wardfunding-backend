package main

import (
	"log"
	"wardfunding/auth"
	"wardfunding/handler"
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
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")
	
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}

// ambil nilai header authorization
// dari header authorization, ambil tokennya saja
// validasi token
// ambil user_id
// ambil user dari db berdasarkan user_id dari service
// set context isinya user 