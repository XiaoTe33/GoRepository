package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/login/:name/:password", func(c *gin.Context) {
		name := c.Param("name")
		password := c.Param("password")
		cookie, err := c.Cookie("cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("cookie", "我有cookie了[doge]", 10, "/", "localhost", false, true)
		}
		c.IndentedJSON(http.StatusOK, gin.H{"name": name, "password": password, "cookie": cookie})
	})

	r.Run(":8080")

}
