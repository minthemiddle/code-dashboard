package main

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// GetUser returns the user details
func GetUser(c *gin.Context) {
	var u User
	DB.C("users").Find(bson.M{"email": jwt.ExtractClaims(c)["id"]}).One(&u)
	c.JSON(200, u)
}
