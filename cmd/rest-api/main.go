package main

import (
	"github.com/echenim/corelightfx/di"
	"github.com/echenim/corelightfx/utilities"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title API
// @version 2.5
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	DbStoreLinker := utilities.Instanciate()
	c, _ := DbStoreLinker.DB()
	defer c.Close()

	storeAPI := di.InitStockAPI(DbStoreLinker)
	nasdaqAPI := di.InitNasdaqAPI(DbStoreLinker)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/stocks", storeAPI.GetAll)
	r.GET("/stocks/findbyname?name=d", storeAPI.FindByName)
	r.GET("/stocks/search?name=d&duration=t", storeAPI.SearchNameAndDuration)
	r.GET("/nasdaq", nasdaqAPI.GetAll)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8082")

}
