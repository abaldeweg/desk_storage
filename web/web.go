package web

import (
	"baldeweg/mission/storage"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func init() {
    log.SetPrefix("web: ")
    log.SetFlags(0)
}

type File struct {
    Name string `json:"name"`
    Body string `json:"body"`
}

type Msg struct {
    Msg string `json:"msg"`
}

func Web() {
    http.HandleFunc("/api/show", makeHandler(showHandler, "GET"))
    http.HandleFunc("/api/update", makeHandler(updateHandler, "PUT"))

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func showHandler(w http.ResponseWriter, r *http.Request) {
    name := filepath.Base(r.URL.Query().Get("name"))

    if !storage.Exists(name) && len(name) >= 3 {
        http.NotFound(w, r)
            return
    }

    c := string(storage.Read(name))
    io.WriteString(w, c)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }

    file := unmarshalJson(string(body))

    if !storage.Exists(file.Name) && len(file.Name) >= 3 {
        http.NotFound(w, r)
            return
    }

    storage.Write(file.Name, file.Body)

    c := string(marshalJson(Msg{Msg: "SUCCESS"}))
    io.WriteString(w, c)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request), method string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != method && r.Method != "OPTIONS" {
            http.NotFound(w, r)
            return
        }

        w.Header().Set("Access-Control-Allow-Origin", getOrigin(r.Header.Get("Origin")))
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
        w.Header().Set("Content-Type", "application/json")

        if r.Method == "OPTIONS" {
            return
        }

        if !isAuthenticated(r.Header.Get("Authorization")) {
            w.WriteHeader(401)
            return
        }

        fn(w, r)
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
