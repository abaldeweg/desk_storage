package router

import (
	"os"

	"github.com/abaldeweg/storage/controller/call"
	"github.com/abaldeweg/storage/controller/mission"
	"github.com/abaldeweg/storage/controller/schedule"
	"github.com/abaldeweg/storage/controller/shift"
	"github.com/abaldeweg/storage/controller/staff"
	"github.com/gin-gonic/gin"
)

func Router() {
    r := gin.New()
    r.SetTrustedProxies(nil)

    if os.Getenv("ENV") != "prod" {
        r.Use(gin.Logger())
    }

	r.Use(gin.Recovery())

    r.Use(headers())

    auth := r.Group("/api", checkAuth)

    // mission
	auth.GET("/mission/show", mission.Show)
	auth.POST("/mission/create", mission.Create)
	auth.PUT("/mission/update", mission.Update)
	auth.GET("/mission/export/html", mission.HtmlExport)

    // shift
    auth.GET("/shift/list", shift.List)

    // staff
    auth.GET("/staff/list", staff.List)

    // schedule
    auth.GET("/schedule/show", schedule.Show)
    auth.PUT("/schedule/update", schedule.Update)

    // call
    auth.GET("/call/list", call.List)
    auth.PUT("/call/update", call.Update)

	r.Run(":8080")
}
