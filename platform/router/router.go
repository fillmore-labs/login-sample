package router

import (
	"encoding/gob"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"login-sample/platform/authenticator"
	"login-sample/platform/middleware"
	"login-sample/web/app/callback"
	"login-sample/web/app/home"
	"login-sample/web/app/login"
	"login-sample/web/app/logout"
	"login-sample/web/app/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	kodata := os.Getenv("KO_DATA_PATH")
	router.Static("/public", kodata+"/static")
	router.LoadHTMLGlob(kodata + "/template/*")

	router.GET("/", home.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
