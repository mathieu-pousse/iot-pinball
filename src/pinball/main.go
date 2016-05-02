package main

import (
	"flag"
	"log"
)

type GlobalConfiguration struct {
	File string
}

var global = GlobalConfiguration{}

func main() {
	log.Println("Starting pinball...")
	flag.StringVar(&global.File, "configuration", "conf/table.json", "set the configuration file")
	flag.Parse()
	CheckSystem()
	LoadConfiguration(global.File)
	table.init()
}
