package mission

import (
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/abaldeweg/storage/controller"
	"github.com/abaldeweg/storage/export/html"
	"github.com/abaldeweg/storage/mission/create"
	"github.com/abaldeweg/storage/storage"
)

func init() {
    log.SetPrefix("web: ")
    log.SetFlags(0)
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {
    filename := "missions.json"
    name := filepath.Base(filename)

    if !storage.Exists(name) && len(name) >= 3 {
        http.NotFound(w, r)
            return
    }

    c := string(storage.Read(name))
    io.WriteString(w, c)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
    filename := "missions.json"

    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }

    file := controller.UnmarshalJson(string(body))

    if len(filename) >= 3 {
        http.NotFound(w, r)
            return
    }

    storage.Write(filename, file.Body)

    c := string(controller.MarshalJson(controller.Msg{Msg: "SUCCESS"}))
    io.WriteString(w, c)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
    create.Create()

    c := string(controller.MarshalJson(controller.Msg{Msg: "SUCCESS"}))
    io.WriteString(w, c)
}

func HtmlExportHandler(w http.ResponseWriter, r *http.Request) {
    c := string(controller.MarshalJson(html.Export()))
    io.WriteString(w, c)
}
