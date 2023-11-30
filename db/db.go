package db

import (
	"log"

	"github.com/raihaninfo/attendance_magagment/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://uc8j6b9b7pjkmbgcuswp:oMjLYeottA8yZ6aErX5PriUKjhxKuf@bqdkti7snuksoswq37kj-postgresql.services.clever-cloud.com/bqdkti7snuksoswq37kj"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.User{}, &models.Class{}, &models.Student{}, &models.Attendance{})

	return db
}
