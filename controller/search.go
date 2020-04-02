package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otz1/otz/service"
)

// Search ...
func (ctx *Controller) Search(c *gin.Context) {
	query := c.Query("query")

	log.Println("Search request for", query)

	ss := service.NewSearchService()
	resp := ss.Search(query)
	c.JSON(http.StatusOK, resp)
}
