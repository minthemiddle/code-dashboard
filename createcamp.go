package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// CreateCamp creates a new Camp
func CreateCamp(c *gin.Context) {
	var camp Camp
	c.BindJSON(&camp)
	if IsAdmin(c) {
		if DB.C("camps").Insert(camp) != nil {
			c.JSON(500, gin.H{"message": "Failed to insert Camp!"})
			return
		}

		body := strings.NewReader(`token=` + os.Getenv("SLACK_TOKEN") + `&name=` + slugify(camp.City) + "17-general")
		req, err := http.NewRequest("POST", "https://slack.com/api/channels.create", body)
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to create channel! Added camp, tho"})
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to create channel!"})
			return
		}
		defer resp.Body.Close()
		var ok interface{}
		json.NewDecoder(resp.Body).Decode(&ok)
		if resp.StatusCode != 200 {
			c.JSON(500, gin.H{"message": "Failed to create channel.", "data": ok})
			return
		}
		c.JSON(200, camp)
	} else {
		c.JSON(403, gin.H{"message": "You are not an Administrator!"})
		return
	}
}
