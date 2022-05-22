package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var secret = []byte("secret")

func ServeApi() {
	router := gin.Default()

	router.Use(sessions.Sessions("session", cookie.NewStore(secret)))

	public := router.Group("/")
	PublicRoutes(public)

	private := router.Group("/")
	private.Use(AuthRequired)
	PrivateRoutes(private)

	router.Run("localhost:8080")
}
