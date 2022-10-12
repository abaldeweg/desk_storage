package router

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func checkAuth(c *gin.Context) {
    if !isAuthenticated(c.GetHeader("Authorization")) {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }
}

func isAuthenticated(auth string) bool {
    token := strings.Split(auth, " ")
    if len(token) == 2 {
        if _, err := checkToken(token[1]); err == nil {
            return true
        }
    }

    return false
}

func checkToken(idToken string) (*auth.Token, error) {
    ctx := context.Background()

    app, err := firebase.NewApp(ctx, nil)
    if err != nil {
        return nil, err
    }

    client, err := app.Auth(ctx)
    if err != nil {
        return nil, err
    }

    token, err := client.VerifyIDTokenAndCheckRevoked(ctx, idToken)
    if err != nil {
        return nil, err
    }

    return token, nil
}
