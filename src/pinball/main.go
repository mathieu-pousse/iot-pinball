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
	log.Println("LED0...")
	table.init()
	log.Println("i2c...")
	table.i2c()
	log.Println("let's read i2c...")
	table.loop()
}
