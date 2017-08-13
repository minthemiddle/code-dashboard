package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"github.com/appleboy/gin-jwt"
)

func DeleteUser(c *gin.Context) {
	if DB.C("users").Remove(bson.M{"email": jwt.ExtractClaims(c)["id"]}) != nil {
		c.JSON(500, gin.H{"message": "Failed to delete user"})
		return
	}
	c.JSON(200, gin.H{"message": "Success!"})
}
