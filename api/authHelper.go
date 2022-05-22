package api

import (
	"fundingSystem/common"
	"fundingSystem/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		log.Println("User not logged in")
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}
func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json common.LoginDto
		session := sessions.Default(c)

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if emptyUserPass(json.User, json.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"status": "\"Parameters can't be empty\""})
			return
		}

		id, err := db.LoginUser(json)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		session.Set(gin.AuthUserKey, id)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func emptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}
