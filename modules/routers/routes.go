package routers

import (
	"howToUseMod/modules/users"
	"howToUseMod/modules/users/UserController"
	"log"
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
		user := UserController.Show(1)
		c.String(http.StatusOK, user.Name)
	})
	userRoutes.POST("/create", func(c *gin.Context) {
		var user users.User
		user.Name = c.PostForm("name")
		user.Email = c.PostForm("email")
		err := user.Create()
		if err != nil {
			log.Fatalln(err)
		}
		c.String(http.StatusOK, "ユーザーの作成に成功しました")
	})
	r.Run()
}
