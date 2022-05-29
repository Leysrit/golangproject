package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfund?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	userByEmail, err := userRepository.FindByEmail("john@gmail.co")

	if err != nil {
		panic(err)
	}

	if userByEmail.ID == 0 {
		fmt.Println("User Not Found")
	} else {
		fmt.Println(userByEmail.Name)
	}

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/user", userHandler.RegisterUser)
	api.POST("/session", userHandler.Login)

	router.Run(":8080")

}