package main

import (
	"crm-service/modules/account"
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
	accountHandler := account.DefaultRequestHandler(db)
	r.POST("/customers", customersHandler.Create)
	r.GET("/customers", accountHandler.AuthMiddleware, customersHandler.Read)
	r.GET("/customersquery", customersHandler.GetByEmail)
	r.GET("/customers/:id", customersHandler.ReadByPk)
	r.PUT("/customers/:id", customersHandler.Update)
	r.DELETE("/customers/:id", customersHandler.Delete)

	r.POST("/register", accountHandler.Create)
	r.GET("/register", accountHandler.ReadByUsername)
	r.POST("/login", accountHandler.Login)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
