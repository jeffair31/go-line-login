package route

import (
	"example/go-test-jwt/src/authen"
	"example/go-test-jwt/src/books"
	"example/go-test-jwt/src/linex"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route() {
	r := gin.New()

	r.Group(``)
	{
		r.POST("/login", authen.LoginHandler)
		r.POST("/line-login", linex.CallbackHandler)

		// Define the callback route
		r.GET("/callback", linex.CallbackHandler)

	}

	rGroup := r.Group("/", authen.Authorizationx)
	{
		rGroup.GET("/books", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
		//
		rGroup.POST("/books", books.CreateBooks)
	}

	r.Run()
}
