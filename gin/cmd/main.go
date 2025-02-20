package main

import (
	"gin/conf"
	"gin/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(":" + conf.HttpPort)
}
