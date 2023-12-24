package auth

import (
	"github.com/gin-gonic/gin"
	"main/models/response"
	"net/http"
	"strings"
)

func ValidateToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, response.Error{Message: "unauthorized"})
		c.Abort()
		return
	}

	values := strings.Split(token, "Bearer ")
	if len(values) != 2 {
		c.JSON(http.StatusUnauthorized, response.Error{Message: "token is not set properly"})
		c.Abort()
		return
	}

	c.Set("email", values[1])
	c.Next()
}
