package main

import (
	"github.com/echenim/corelightfx/di"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=nopassword dbname=store-store port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io:8080
// @BasePath /v2
func main() {
	cn := initDb()
	c, _ := cn.DB()
	defer c.Close()

	storeAPI := di.InitStockAPI(cn)

	r := gin.Default()
	r.GET("/products", storeAPI.FindAll)

	er := r.Run()
	if er != nil {
		panic(er)
	}
}
