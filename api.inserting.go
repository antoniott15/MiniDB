package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (api *API) createNewRecord(r *gin.RouterGroup) {
	r.POST("/insert-record/:table", func(c *gin.Context) {
		table := c.Param("table")

		items := make(map[string]interface{})

		if err := c.ShouldBind(&items); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		key, err := strconv.Atoi(fmt.Sprint(items["KEY"]))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		t := api.engine.getTableByName(table)
		str := &Structure{
			Key:     key,
			Headers: t.Headers,
			Attribs: items,
		}
		if err := api.engine.insertIntoTable(str, table); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logrus.Infof("Saving new record: \n%+v\n", str)
		c.JSON(http.StatusOK, gin.H{"data": gin.H{
			"key": str.Key,
			"headers": str.Headers,
			"record": str.Attribs,
		}})
		return
	})

}
