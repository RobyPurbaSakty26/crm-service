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
	dsn := "root:password@tcp(db:3306)/mini_project?parseTime=true"
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
	r.POST("/customers", accountHandler.AuthMiddleware, customersHandler.Create)
	r.GET("/customers", accountHandler.AuthMiddleware, customersHandler.Read)
	r.GET("/customers/:id", accountHandler.AuthMiddleware, customersHandler.ReadByPk)
	r.PUT("/customers/:id", customersHandler.Update)
	r.DELETE("/customers/:id", accountHandler.AuthMiddleware, customersHandler.Delete)

	r.POST("/register", accountHandler.Create)
	r.GET("/actors", accountHandler.AuthMiddleware, accountHandler.AuthMiddleware, accountHandler.ReadByUsername)
	r.POST("/login", accountHandler.Login)
	r.PUT("/approval/:id", accountHandler.AuthMiddleware, accountHandler.Update)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
