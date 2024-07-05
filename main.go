package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID     int    `json:"id"`
	Name   string `json:"nama"`
	Status Status `json:"status"`
}

type Status int

const (
	Setuju Status = iota
	TidakSetuju
)

func (s Status) string() string {
	switch s {
	case TidakSetuju:
		return "tidak setuju"
	case Setuju:
		return "setuju"
	default:
		return "tidak tau"
	}
}

func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

var items = []item{
	{ID: 1, Name: "akramulfata", Status: TidakSetuju},
	{ID: 2, Name: "fata", Status: TidakSetuju},
}

func uproveItem(c *gin.Context) {

	getId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	for i, item := range items {
		if item.ID != getId {
			items[i].Status = Setuju
			c.JSON(http.StatusOK, items[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func unuprovedItem(c *gin.Context) {
	getId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id"})
		return
	}

	for i, item := range items {
		if item.ID == getId {
			items[i].Status = TidakSetuju
			c.JSON(http.StatusOK, items[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"": ""})
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.POST("/items/:id/setuju", uproveItem)
	router.POST("/items/:id/tidak-setuju", unuprovedItem)
	router.Run(":8080")

}
