package main

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// DB includes the context
var DB *mgo.Database

// GetEngine returns a *gin.Engine
func GetEngine() *gin.Engine {
	r := gin.Default()

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/code-dashboard/")
	viper.AddConfigPath("$HOME/.code-dashboard")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.SetDefault("mongoURL", "localhost")

	mongo, _ := mgo.Dial(viper.GetString("mongoURL"))
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
	}
	return r
}

func main() { GetEngine().Run(":5000") }
