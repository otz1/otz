package main

import (
	"github.com/otz1/otz/cache"
	"github.com/otz1/otz/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ctrl := controller.New()
	router := gin.Default()

	{
		conf := cors.Default()
		router.Use(conf)
	}

	router.GET("/search", ctrl.Search)
	router.Run(":8001")
	cache.Close()
}
