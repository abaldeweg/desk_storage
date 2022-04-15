package html

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/abaldeweg/storage/mission/create"
)

func TestConfig(t *testing.T) {
    old := file
    defer func() { file = old }()

    file = func(string) []byte {
        d, err := json.Marshal(create.Logfile{Missions: []create.Mission{}})
        if err != nil {
            t.Fatal(err)
        }

        return d
    }

	export := Export()
    if reflect.TypeOf(export).String() != "string" {
        t.Fatalf("%s is not type of %s", reflect.TypeOf(export),  "string")
    }
}
