package main

import (
	"flag"
	"mongo-event-cacher/cmd/app"
	"mongo-event-cacher/config"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func main() {
	flag.Parse()
	cfg := config.NewConfig(*configFlag)
	app.NewListener(cfg)
}
