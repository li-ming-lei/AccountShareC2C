package main

//created by H.G Nuwan Indika

import (
	"github.com/gin-gonic/gin"
	"github.com/li-ming-lei/AccountShareC2C/controllers"
	"net/http"
)

func main() {
	// Get a UserController instance
	// Get a user resource
	router := gin.Default()
	router.Use(Cors())
	router.OPTIONS("/", controllers.OptionsRequest)
	router.POST("/v1/want", controllers.AddWant)
	router.DELETE("/v1/want", controllers.DeleteWant)
	router.POST("/v1/listwants", controllers.ListWants)
	router.POST("/v1/forsell", controllers.AddForSell)
	router.POST("/v1/listforsells", controllers.ListForSells)
	router.DELETE("/v1/forsell", controllers.DeleteForSell)

	router.POST("/v1/get-user", controllers.GetUser)

	router.Run(":8000")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
