package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func (api *API) createTable(r *gin.RouterGroup) {
	r.POST("/create-table", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "wa"})
		return
	})

}
