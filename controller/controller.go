package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type File struct {
    Body string `json:"body"`
}

type Msg struct {
    Msg string `json:"msg"`
}

func init() {
    log.SetPrefix("controller: ")
    log.SetFlags(0)
}

func Cors() gin.HandlerFunc {
    return cors.New(cors.Config{
        AllowOrigins: strings.Split(os.Getenv("CORS_ALLOW_ORIGIN"), ","),
        AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
        AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders: []string{"Content-Length"},
    })
}

func Auth(c *gin.Context) {
    if !isAuthenticated(c.GetHeader("Authorization")) {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }
}

func UnmarshalJson(blob []byte, d *[]interface{}) *[]interface{} {
	if err := json.Unmarshal(blob, &d); err != nil {
		log.Fatal(err)
	}

    return d
}

func MarshalJson(data interface{}) []byte {
	d, err := json.Marshal(&data)
    if err != nil {
        log.Fatal(err)
    }

    return d
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
