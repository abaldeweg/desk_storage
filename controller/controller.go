package controller

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type File struct {
    Body string `json:"body"`
}

func init() {
    log.SetPrefix("web: ")
    log.SetFlags(0)
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

func unmarshalJson(blob string) File {
    var d File
	if err := json.Unmarshal([]byte(blob), &d); err != nil {
		log.Fatal(err)
	}

    return d
}

func marshalJson(data interface{}) []byte {
	d, err := json.Marshal(&data)
    if err != nil {
        log.Fatal(err)
    }

    return d
}

func getOrigin(origin string) string {
    hosts := strings.Split(os.Getenv("CORS_ALLOW_ORIGIN"), ",")

    if inSlice(origin, hosts) {
        return origin
    }

    return "null"
}

func inSlice(origin string, list []string) bool {
    for _, item := range list {
        if item == origin {
            return true
        }
    }

    return false
}
