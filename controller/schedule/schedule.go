package schedule

import (
	"log"

	"github.com/abaldeweg/storage/storage"
	"github.com/gin-gonic/gin"
)

var filename = "schedule.json"

func init() {
    log.SetPrefix("schedule: ")
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
