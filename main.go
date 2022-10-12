package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abaldeweg/storage/controller/call"
	"github.com/abaldeweg/storage/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    if _, err := os.Stat("./.env"); err == nil {
        if err := godotenv.Load(); err != nil {
            log.Fatal("Error loading .env file")
        }
    }

    gin.SetMode(getGinMode())

    fmt.Println("baldeweg/storage <https://github.com/abaldeweg/storage>")

    action := "web"
    if len(os.Args) == 2 {
        if os.Args[1] == "cron" {
            action = "cron"
        }
    }

    if action == "cron" {
        call.Set()
    } else {
        router.Router()
    }
}

func getGinMode() string {
    mode := os.Getenv("ENV")

    switch mode {
    case "prod":
        return gin.ReleaseMode
    case "dev":
        return gin.DebugMode
    case "test":
        return gin.TestMode
    default:
        return gin.ReleaseMode
    }
}
