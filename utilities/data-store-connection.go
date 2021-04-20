package utilities

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Instanciate() *gorm.DB {
	dsn := "host=dev-corelight user=postgres password=nopassword dbname=store-store port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=nopassword dbname=store-store port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
