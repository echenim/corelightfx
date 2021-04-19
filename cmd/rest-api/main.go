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
