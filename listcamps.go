package main

import "github.com/gin-gonic/gin"

func ListCamps(c *gin.Context) {
	var camps []Camp
	if DB.C("camps").Find(nil).All(&camps) != nil {
		c.JSON(500, gin.H{"message": "Failed to get camps"})
		return
	}
	c.JSON(200, camps)
}
