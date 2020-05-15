package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (api *API) createTable(r *gin.RouterGroup) {
	r.POST("/create-table", func(c *gin.Context) {
		items := make(map[string]string)
		if err := c.ShouldBind(&items); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if items["KEY"] == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Table need a key"})
			return
		}
		var headers []string
		for k, _ := range items {
			if strings.ToUpper(k) != "TABLE" {
				headers = append(headers,strings.ToUpper(k))
			}
		}
		table, err := api.engine.createNewTable(items["TABLE"], headers)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		logrus.Infof("Create new table with name: %s", items["TABLE"])

		c.JSON(http.StatusOK, gin.H{"data": table})
		return
	})

}
