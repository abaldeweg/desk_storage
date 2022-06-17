package staff

import (
	"encoding/json"
	"log"

	"github.com/abaldeweg/storage/storage"
	"github.com/gin-gonic/gin"
)

var filename = "staff.json"

type Staff struct {
    Key string `json:"key"`
    Phone string `json:"phone"`
    Value string `json:"value"`
}

func init() {
    log.SetPrefix("staff: ")
    log.SetFlags(0)
}

func List(c *gin.Context) {
    if !storage.Exists(filename) {
        c.AbortWithStatus(404)
        return
    }

    var d []Staff
    if err := json.Unmarshal(storage.Read(filename), &d); err != nil {
        c.AbortWithStatus(404)
        return
    }

    c.JSON(200, d)
}
