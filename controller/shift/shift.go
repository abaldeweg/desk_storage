package shift

import (
	"encoding/json"
	"log"

	"github.com/abaldeweg/storage/storage"
	"github.com/gin-gonic/gin"
)

var filename = "shift.json"

type Shift struct {
    Name string `json:"name"`
    Desc string `json:"desc"`
    Start string `json:"start"`
    Duration int `json:"duration"`
}

func init() {
    log.SetPrefix("shift: ")
    log.SetFlags(0)
}

func List(c *gin.Context) {
    if !storage.Exists(filename) {
        c.AbortWithStatus(404)
        return
    }

    var d []Shift
    if err := json.Unmarshal(storage.Read(filename), &d); err != nil {
        c.AbortWithStatus(404)
        return
    }

    c.JSON(200, d)
}
