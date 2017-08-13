package main

import (
	"github.com/gin-gonic/gin"
)

func CreateCamp(c *gin.Context) {
	var camp Camp
	c.BindJSON(&camp)
	if IsAdmin(c) {
		if DB.C("camps").Insert(camp) != nil {
			c.JSON(500, gin.H{"message": "Failed to insert Camp!"})
			return
		}
		c.JSON(200, camp)
	} else {
		c.JSON(403, gin.H{"message": "You are not an Administrator!"})
		return
	}
}
