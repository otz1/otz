package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/otz1/otz/cache"
	"github.com/otz1/otz/controller"
	"github.com/otz1/otz/util"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	env := util.GetEnv("ENVIRONMENT", "local")
	log.Println("RUNNING IN", strings.ToUpper(env))

	err := sentry.Init(sentry.ClientOptions{
		Dsn:         "https://b8dd734fc88c43eb95110d9a1a06a2e1@sentry.io/5187014",
		Environment: env,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	ctrl := controller.New()
	router := gin.Default()
	{
		conf := cors.Default()
		router.Use(conf)
	}

	router.GET("/loaderio-f9e388176893839315d3c6612a2739b0.txt", func(c *gin.Context) {
		c.String(http.StatusOK, "loaderio-f9e388176893839315d3c6612a2739b0")
	})

	router.GET("/search", ctrl.Search)
	router.Run(":8001")
	cache.Close()
}
