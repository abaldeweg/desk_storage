package schedule

import (
	"encoding/json"
	"log"

	"github.com/abaldeweg/storage/controller"
	"github.com/abaldeweg/storage/storage"
	"github.com/gin-gonic/gin"
)

var filename = "schedule.json"

type Schedule struct {
    Staff string `json:"staff"`
    Starttime string `json:"starttime"`
    Endtime string `json:"endtime"`
}

type Request struct {
    Body interface{} `json:"body"`
}

func init() {
    log.SetPrefix("schedule: ")
    log.SetFlags(0)
}

func Show(c *gin.Context) {
    if !storage.Exists(filename) {
        c.AbortWithStatus(404)
        return
    }

    var d map[string][]Schedule
    if err := json.Unmarshal(storage.Read(filename), &d); err!= nil {
        c.AbortWithStatus(404)
        return
    }

    c.JSON(200, d)
}

func Update(c *gin.Context) {
    var file Request
    if err := c.ShouldBind(&file); err != nil {
        log.Fatal(err)
        c.AbortWithStatus(404)
        return
    }

    storage.Write(filename, string(controller.MarshalJson(file.Body)))

    d := controller.Msg{Msg: "SUCCESS"}

    c.JSON(200, d)
}
