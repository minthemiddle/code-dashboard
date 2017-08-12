package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type register struct {
	Email string `json:"email" binding:"required"`
	Pass  string `json:"password" binding:"required"`
}

// RegisterHandler registers a new user
func RegisterHandler(c *gin.Context) {
	var r register
	c.BindJSON(&r)
	if n, _ := DB.C("users").Find(bson.M{"email": r.Email}).Count(); n == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(r.Pass), bcrypt.DefaultCost)
		a := User{Email: r.Email, Password: string(hashedPassword[:])}
		if DB.C("users").Insert(a) != nil {
			c.JSON(500, gin.H{"message": "Failed to save user"})
			return
		}
		c.JSON(200, gin.H{"message": "success!"})
	} else {
		c.JSON(400, gin.H{"message": "User already exists"})
	}
}
