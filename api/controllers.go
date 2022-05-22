package api

import (
	"fmt"
	"fundingSystem/common"
	"fundingSystem/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json common.UserDto

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		session := sessions.Default(c)
		id, _ := db.CreateUser(json)
		session.Set(gin.AuthUserKey, id)
	}
}

func CreateProjectHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json common.ProjectDto

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		session := sessions.Default(c)
		userId := session.Get(gin.AuthUserKey).(int)
		id, _ := db.CreateProject(json, userId)
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/project?id=%d", id))
	}
}
