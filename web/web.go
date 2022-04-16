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
    // mission
    http.HandleFunc("/api/mission/show", controller.MakeHandler(mission.ShowHandler, "GET"))
    http.HandleFunc("/api/mission/update", controller.MakeHandler(mission.UpdateHandler, "PUT"))
    http.HandleFunc("/api/mission/create", controller.MakeHandler(mission.CreateHandler, "POST"))
    http.HandleFunc("/api/mission/export/html", controller.MakeHandler(mission.HtmlExportHandler, "GET"))

    //  shift

    log.Fatal(http.ListenAndServe(":8080", nil))
}
