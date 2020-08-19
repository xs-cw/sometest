package handler

import (
	"demo-go/response"
	"demo-go/servers/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	req := models.LoginRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(response.Failed(http.StatusInternalServerError, "parse request failed:"+err.Error()))
		return
	}
	if req.UserName == "" {
		req.Password = ""
		// return
	}
	t, err := NewToken()
	if err != nil {
		c.JSON(response.Failed(http.StatusInternalServerError, "create token failed:"+err.Error()))
		return
	}
	c.JSON(response.Success(models.LoginResponse{Token: t}))
}
