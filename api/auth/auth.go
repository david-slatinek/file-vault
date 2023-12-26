package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
	"log"
	"main/models/response"
	"net/http"
	"strings"
	"time"
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

	log.Printf("token: %s", values[1])

	email, err := getEmail(values[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Error{Message: err.Error()})
		c.Abort()
		return
	}

	c.Set("email", email)
	c.Next()
}

func getEmail(token string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	users, _, err := client.Users.ListEmails(ctx, nil)

	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", nil
	}

	for _, user := range users {
		if user.GetPrimary() {
			return user.GetEmail(), nil
		}
	}

	return "", nil
}
