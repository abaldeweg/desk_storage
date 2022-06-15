package call

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
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

func List(c *gin.Context) {
    url := "https://api.sipgate.com/v2/w0/phonelines/p0/forwardings"

    req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
    if err != nil {
		log.Fatal(err)
	}

    c.JSON(200, request(req))
}

func Update(c *gin.Context) {
    url := "https://api.sipgate.com/v2/w0/phonelines/p0/forwardings"

    var file Forwarding
    forwarding := []Forwarding{}

    if err := c.ShouldBind(&file); err == nil {
        forwarding = []Forwarding{{file.Destination, 0, true}}
    }

    data, err := json.Marshal(Forwardings{forwarding})
    if err != nil {
        log.Fatal(err)
	}

    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
    if err != nil {
		log.Fatal(err)
	}

    request(req)

    c.JSON(200, "UPDATED")
}

func request(req *http.Request) string {
    token := []byte(os.Getenv("SIPGATE_TOKEN_NAME") + ":" + os.Getenv("SIPGATE_PAT"))

	req.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString(token))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
		log.Fatal(err)
	}

    return string(body)
}
