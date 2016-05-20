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
	log.Println("init...")
	table.init()
	log.Println("LED0...")
	table.OnBoardLeds()
	log.Println("i2c...")
	table.I2CLeds()
	log.Println("i2c again...")
	table.I2CLeds()
	log.Println("let's read i2c...")
	table.loop()
}
