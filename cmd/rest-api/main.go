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

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api")
	{
		s := v1.Group("/stocks")
		{
			s.GET("", storeAPI.GetAll)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8082")

	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// v1 := r.Group("/api/v1")
	// {
	// 	v1.GET("/stocks", storeAPI.FindAll)
	// 	v1.GET("/stocks/v1", storeAPI.FindAll)
	// 	v1.GET("/stocks/v2", storeAPI.FindAll)
	// }

	// er := r.Run()
	// if er != nil {
	// 	panic(er)
	// }
}
