package main

import (
	"srebootcamp/internal/config"
	"srebootcamp/internal/db"
	"srebootcamp/internal/handler"
	"srebootcamp/internal/router"
)

func main() {
	c := config.LoadConfig()
	dbConn := db.InitDB(c)
	defer dbConn.Close()

	h := &handler.Handler{DB: dbConn}
	r := router.SetupRouter(h)
	_ = r.Run(c.Port)
}
