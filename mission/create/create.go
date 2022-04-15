package create

import (
	"baldeweg/mission/filetypes"
	"baldeweg/mission/storage"
	"encoding/json"
	"log"
	"time"
)

func init() {
    log.SetPrefix("create: ")
    log.SetFlags(0)
}

func Create(){
    filename := "missions.json"

    create := filetypes.Mission{
        Date: time.Now().Format("2006-01-02"),
        Time: time.Now().Format("15:04"),
    }

    var t filetypes.Logfile
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
