package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.POST("/login", LoginPostHandler())
	g.POST("/register", RegisterPostHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/dashboard", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		session := sessions.Default(c)
		user := session.Get(gin.AuthUserKey).(int)
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	})
	g.GET("/project")
	g.POST("/createProject", CreateProjectHandler())

}
