package file

import (
	"log"
	"os"
)

func init() {
    log.SetPrefix("file: ")
    log.SetFlags(0)
}

func Write(filename string, content string) {
    if err := os.WriteFile(os.Getenv("FILE_PATH") + "/" + filename, []byte(content), 0644); err != nil {
        log.Fatal(err)
    }
}

func Read(filename string) []byte {
    data, err := os.ReadFile(os.Getenv("FILE_PATH") + "/" + filename)
    if err != nil {
        log.Fatal(err)
    }

    return data
}

func Exists(filename string) bool {
    if _, err := os.Stat(os.Getenv("FILE_PATH") + "/" + filename); err == nil {
        return true
    }

    return false
}
