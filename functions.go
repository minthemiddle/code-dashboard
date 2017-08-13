package main

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) bool {
	var u User
	if DB.C("users").Find(bson.M{"email": jwt.ExtractClaims(c)["id"]}).One(&u) != nil {
		return false
	}
	return u.Admin
}
