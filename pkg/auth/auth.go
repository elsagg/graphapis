package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var UserCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// X-Endpoint-API-UserInfo

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User

		header := c.Request.Header.Get("X-Endpoint-API-UserInfo")

		if header == "" {
			c.Next()
			return
		}

		if err := json.Unmarshal([]byte(header), &user); err != nil {
			displayError := errors.New("Failed to authenticate user")
			c.AbortWithError(http.StatusForbidden, displayError)
			return
		}

		ctx := context.WithValue(c.Request.Context(), UserCtxKey, user)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func GetUser(ctx context.Context) User {
	raw, _ := ctx.Value(UserCtxKey).(User)
	return raw
}
