package routers

import (
	"howToUseMod/modules/users/UserController"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ホーム")
	})
	userRoutes := r.Group("/users")
	userRoutes.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, UserController.Show(1).Name)
	})
	r.Run()
}
