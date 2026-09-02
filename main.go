package main

import (
	"srebootcamp/internal/config"
	"srebootcamp/internal/router"
)

func main() {
	r := router.SetupRouter()
	c := config.LoadConfig()
	_ = r.Run(c.Port)

}
