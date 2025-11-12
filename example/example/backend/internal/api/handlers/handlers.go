package handlers

import (
	"github.com/gin-gonic/gin"
)

func ListItems(c *gin.Context) {
	c.JSON(200, gin.H{"items": []string{}})
}

func CreateItem(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Item created"})
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id})
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "message": "Item updated"})
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "message": "Item deleted"})
}
