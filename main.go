package main

import (
	"fmt"

	"github.com/raihaninfo/attendance_magagment/controllers"
	"github.com/raihaninfo/attendance_magagment/db"
	_ "github.com/raihaninfo/attendance_magagment/docs"
)

var Port string = ":3000"

func main() {
	DB := db.Init()
	fmt.Println("Listing to port : ", Port)
	controllers.Controller(Port, DB)
	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		user := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message, "user": user})
	})
}
