package call

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Forwardings struct {
    Forwardings []Forwarding `json:"forwardings"`
}

type Forwarding struct {
    Destination string `json:"destination"`
    Timeout int `json:"timeout"`
    Active bool `json:"active"`
}

func init() {
    log.SetPrefix("call: ")
    log.SetFlags(0)
}

func Update(c *gin.Context) {
    var file Forwarding
    if err := c.ShouldBind(&file); err != nil {
        c.AbortWithStatus(404)
        return
    }

    request([]Forwarding{{file.Destination, 0, true}})

    c.JSON(200, "UPDATED")
}

func Reset(c *gin.Context) {
    request([]Forwarding{})

    c.JSON(200, "RESET")
}

func request(forwardings []Forwarding) {
    url := "https://api.sipgate.com/v2/w0/phonelines/p0/forwardings"

    data, err := json.Marshal(Forwardings{forwardings})
    if err != nil {
        log.Fatal(err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
    if err != nil {
		log.Fatal(err)
	}

    token := []byte(os.Getenv("SIPGATE_TOKEN_NAME") + ":" + os.Getenv("SIPGATE_PAT"))

	req.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString(token))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	res.Body.Close()
}
