package create

import (
	"encoding/json"
	"log"
	"time"

	"github.com/abaldeweg/storage/storage"
)

type Logfile struct {
    Notes []string `json:"notes"`
    Replacements map[string]string `json:"replacements"`
    Missions []Mission `json:"missions"`
}

type Mission struct {
    Date string `json:"date"`
    Time string `json:"time"`
    Keyword string `json:"keyword"`
    Situation string `json:"situation"`
    Unit string `json:"unit"`
    Location string `json:"location"`
    Links []string `json:"links"`
    Private bool `json:"private"`
}

func init() {
    log.SetPrefix("create: ")
    log.SetFlags(0)
}

func Create(){
    filename := "missions.json"

    create := Mission{
        Date: time.Now().Format("2006-01-02"),
        Time: time.Now().Format("15:04"),
    }

    var t Logfile
	if err := json.Unmarshal([]byte(string(storage.Read(filename))), &t); err != nil {
		log.Fatal(err)
	}

    t.Missions = append(t.Missions, create)


    d, err := json.Marshal(&t)
    if err != nil {
        log.Fatal(err)
    }

    storage.Write(filename, string(d))
}
