package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/otz1/otz/cache"
	"github.com/otz1/otz/controller"
)

func main() {
	ctrl := controller.New()
	router := gin.Default()

	{
		conf := cors.Default()
		router.Use(conf)
	}

	router.GET("/search", ctrl.Search)
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "this is a test!")
	})

	router.Run(":8001")
	cache.Close()
}
