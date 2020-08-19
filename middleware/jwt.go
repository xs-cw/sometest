package middleware

import (
	"fmt"
	"net/http"

	"demo-go/config"
	"demo-go/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func MiddleJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetSecretKey()), nil
		})

		if err != nil {
			c.Abort()
			c.JSON(response.Failed(http.StatusInternalServerError, "parse request failed:"+err.Error()))
			return
		}
		if !t.Valid {
			c.Abort()
			c.JSON(response.Failed(http.StatusInternalServerError, "token invalid!"))
			return
		}
		fmt.Println(t)
		c.Next()
	}
}
