package staff

import (
	"log"

	"github.com/abaldeweg/storage/storage"
	"github.com/gin-gonic/gin"
)

var filename = "staff.json"

func init() {
    log.SetPrefix("staff: ")
    log.SetFlags(0)
}

func List(c *gin.Context) {
    if !storage.Exists(filename) {
        c.AbortWithStatus(404)
        return
    }

    d := string(storage.Read(filename))

    c.JSON(200, d)
}
