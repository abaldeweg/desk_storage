package main

import (
	"baldeweg/mission/web"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
    log.SetPrefix("main: ")
    log.SetFlags(0)
}

func main() {
    if _, err := os.Stat("./.env"); err == nil {
        if err := godotenv.Load(); err != nil {
            log.Fatal("Error loading .env file")
        }
    }

    fmt.Println("baldeweg/storage <https://github.com/abaldeweg/storage>")

    web.Web()
}
