package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otz1/otz/service"
)

var searchService = service.NewSearchService()

// Search ...
func (ctx *Controller) Search(c *gin.Context) {
	query := c.Query("query")
	resp := searchService.Search(query)
	c.PureJSON(http.StatusOK, resp)
}
