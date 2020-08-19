package router

import (
	"demo-go/middleware"
	"demo-go/servers/handler"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	g := e.Group("/v1")
	g.Use(middleware.Cors())
	g.GET("/users/login", handler.Login)
	g.Static("/static", "./static")
	g.Use(middleware.MiddleJWTAuth())
	g.GET("/getlist", handler.GetList)

}
