package main

import (
	"strings"
	"time"

	"github.com/itsjamie/gin-cors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// DB includes the context
var DB *mgo.Database

// GetEngine returns a *gin.Engine
func main() {
	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	mongo, mongoerr := mgo.Dial("localhost")
	if mongoerr != nil {
		fmt.Println("Could not connect to db")
		return
	}
	DB = mongo.DB("code-dashboard")

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "code-dashboard",
		Key:        []byte("apfelstrudel"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if n, err := DB.C("users").Find(bson.M{"email": userId, "password": password}).Count(); err != nil && n >= 1 {
				return userId, false
			}
			return userId, true
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			if DB.C("users").Find(bson.M{"email": userId}).One(nil) != nil {
				return false
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", RegisterHandler)

	auth := r.Group("/api/v1")

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// refresh your token
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

		// Get current user
		auth.GET("/user", GetUser)

		// Change a user
		auth.PUT("/user", ChangeUser)

		// Delete a user
		auth.DELETE("/user", DeleteUser)

		// Get all camps
		auth.GET("/camps", ListCamps)

		// Create a camp
		auth.POST("/camps", CreateCamp)
	}
	r.Run(":3000")
}

func slugify(s string) string { return strings.Replace(s, " ", "-", -1) }
