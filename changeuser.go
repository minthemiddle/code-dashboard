package main

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// ChangeUser changes a User
func ChangeUser(c *gin.Context) {
	var u LimitedUser
	c.BindJSON(&u)
	var user User
	if DB.C("users").Find(bson.M{"email": jwt.ExtractClaims(c)["id"]}).One(&user) != nil {
		c.JSON(500, gin.H{"message": "Failed to get user!"})
		return
	}

	//TODO: Refactoring
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Parent = u.Parent

	if DB.C("users").Update(bson.M{"email": jwt.ExtractClaims(c)["id"]}, user) != nil {
		c.JSON(500, gin.H{"message": "Failed to save user!"})
		return
	}
	c.JSON(200, user)
}
