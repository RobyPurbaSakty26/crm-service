package main

import (
	"crm-service/modules/customers"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	dsn := "root:1234@tcp(localhost:3306)/mini_project?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	customersHandler := customers.DefaultRequestHandler(db)
	r.POST("/customers", customersHandler.Create)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
