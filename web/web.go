package web

import (
	"log"
	"net/http"

	"github.com/abaldeweg/storage/controller"
	"github.com/abaldeweg/storage/controller/mission"
)

func init() {
    log.SetPrefix("web: ")
    log.SetFlags(0)
}

func Web() {
    http.HandleFunc("/api/show", controller.MakeHandler(mission.ShowHandler, "GET"))
    http.HandleFunc("/api/update", controller.MakeHandler(mission.UpdateHandler, "PUT"))
    http.HandleFunc("/api/create", controller.MakeHandler(mission.CreateHandler, "POST"))
    http.HandleFunc("/api/export/html", controller.MakeHandler(mission.HtmlExportHandler, "GET"))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
