package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		cookie := "NotSet"
		fmt.Println("当前的cookie为:", cookie)
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
	c.String(http.StatusOK, "cookie success!")
}
