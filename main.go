package main

import (
	"flag"

	"github.com/brickstudy/blockchain-module/src/app"
	"github.com/brickstudy/blockchain-module/src/config"
)

var (
	configFlag = flag.String("env", "./environment.toml", "Not found environment toml file")
	difficulty = flag.Int("difficulty", 12, "Error difficulty")
)

func main() {
	flag.Parse()

	c := config.NewConfig(*configFlag)
	app.NewApp(c, int64(*difficulty))
}
