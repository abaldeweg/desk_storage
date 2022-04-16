package mission

import (
	"io"
	"log"

	"github.com/abaldeweg/storage/controller"
	"github.com/abaldeweg/storage/export/html"
	"github.com/abaldeweg/storage/mission/create"
	"github.com/abaldeweg/storage/storage"
	"github.com/gin-gonic/gin"
)

var filename = "missions.json"

func init() {
    log.SetPrefix("mission: ")
    log.SetFlags(0)
}

func Show(c *gin.Context) {
    if !storage.Exists(filename) {
        c.AbortWithStatus(404)
        return
    }

    d := string(storage.Read(filename))

    c.JSON(200, d)
}

func Create(c *gin.Context) {
    create.Create()

    d := controller.Msg{Msg: "SUCCESS"}

    c.JSON(200, d)
}

func Update(c *gin.Context) {
    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.AbortWithStatus(404)
        return
    }

    file := controller.UnmarshalJson(string(body))
    storage.Write(filename, file.Body)

    d := controller.Msg{Msg: "SUCCESS"}

    c.JSON(200, d)
}

func HtmlExport(c *gin.Context) {
    d := html.Export()

    c.JSON(200, d)
}
