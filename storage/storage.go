package storage

import (
	"baldeweg/mission/storage/file"
	"baldeweg/mission/storage/gcpBucket"
	"log"
	"os"
)

type Adapter struct {
    Read func(string) []byte
    Write func(string, string)
    Exists func(string) bool
}

var Adapters = map[string]Adapter{
    "file":{file.Read, file.Write, file.Exists},
    "gcp-bucket":{gcpBucket.Read, gcpBucket.Write, gcpBucket.Exists},
}

func init() {
    log.SetPrefix("storage: ")
    log.SetFlags(0)
}

func Write(filename string, content string) {
    func(fn func(string, string), filename string, content string)  {
        fn(filename, content)
        }(Adapters[os.Getenv("STORAGE")].Write, filename, content)
    }

func Read(filename string) []byte {
    return func(fn func(string) []byte, filename string) []byte {
        return fn(filename)
    }(Adapters[os.Getenv("STORAGE")].Read, filename)
}

func Exists(filename string) bool {
    return func(fn func(string) bool, filename string) bool {
        return fn(filename)
    }(Adapters[os.Getenv("STORAGE")].Exists, filename)
}
