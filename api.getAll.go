package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (api *API) GetAll(r *gin.RouterGroup) {
	r.GET("/records/:table", func(c *gin.Context) {
		table := c.Param("table")

		tables := api.engine.getTableByName(table)
		headers := tables.Headers

		keys := tables.StructureKeys
		var records []*Record

		for _, elements := range keys {
			re, _ := tables.StructTree.Search(elements)
			records = append(records, re)
		}

		var Structs []*Structure
		for _, elements := range records {
			var Struct *Structure
			 _ = json.Unmarshal(elements.Value, &Struct)
			 Structs = append(Structs, Struct)
		}

		var record []map[string]interface{}

		for _,elements := range Structs {
			record = append(record, elements.Attribs)
		}

		logrus.Infof("Sending headers \n %+v \n and records \n%+v\n", headers ,record)

		c.JSON(http.StatusOK, gin.H{"data": gin.H{
			"headers":  headers,
			"records": record,
		}})
		return
	})

	r.GET("/tables", func(c *gin.Context) {
		var tableName []string
		for _,elements := range api.engine.TablesTree {
			tableName = append(tableName, elements.Name)
		}

		logrus.Infof("Sending tables: \n%+v\n", tableName)
		c.JSON(http.StatusOK, gin.H{"data": tableName})
		return

	})

}
