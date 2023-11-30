package main

import (
	"fmt"

	"github.com/raihaninfo/attendance_magagment/controllers"
	"github.com/raihaninfo/attendance_magagment/db"
	_ "github.com/raihaninfo/attendance_magagment/docs"
)



func main() {
	DB := db.Init()
	fmt.Println("Listing to port : ", Port)
	controllers.Controller( DB)
}
