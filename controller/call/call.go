package call

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/abaldeweg/storage/controller/schedule"
	"github.com/abaldeweg/storage/controller/staff"
	"github.com/abaldeweg/storage/storage"
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

func Set() {
    schedule, err := findSchedule()
    if err != nil {
        return
    }

    staff, err := findStaff(schedule.Staff)
    if err != nil {
        return
    }

    url := "https://api.sipgate.com/v2/w0/phonelines/p0/forwardings"

    forwarding := []Forwarding{{staff.Phone, 0, true}}

    data, err := json.Marshal(Forwardings{forwarding})
    if err != nil {
        log.Fatal(err)
	}

    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
    if err != nil {
		log.Fatal(err)
	}

    request(req)
}

func findSchedule() (schedule.Schedule, error) {
    var filename = "schedule.json"

    if !storage.Exists(filename) {
        return schedule.Schedule{}, errors.New("file not found")
    }

    var d []schedule.Schedule
    if err := json.Unmarshal(storage.Read(filename), &d); err != nil {
        return schedule.Schedule{}, errors.New("no valid JSON")
    }

    now := time.Now().Unix()

    for _, s := range d {
        if int64(s.Start) <= now && int64(s.End) >= now {
            return s, nil
        }
    }

    return schedule.Schedule{}, errors.New("nothing found")
}

func findStaff(id string) (staff.Staff, error) {
    var filename = "staff.json"

    if !storage.Exists(filename) {
        return staff.Staff{}, errors.New("file not found")
    }

    var d []staff.Staff
    if err := json.Unmarshal(storage.Read(filename), &d); err != nil {
        return staff.Staff{}, errors.New("no valid JSON")
    }

    for _, s := range d {
        if id == s.Id {
            return s, nil
        }
    }

    return staff.Staff{}, errors.New("nothing found")
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
