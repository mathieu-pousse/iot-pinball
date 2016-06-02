package main

import (
	"flag"
	"log"
)

type GlobalConfiguration struct {
	File string
	NoCheck bool
	NoHardware bool
}


var global = GlobalConfiguration{

}


func parseArguments() {
	flag.BoolVar(&global.NoCheck, "no-check", false, "do not perform start up check-up")
	flag.StringVar(&global.File, "configuration", "conf/table.json", "table configuration file")
	flag.BoolVar(&global.NoHardware, "no-hardware", false, "are we running w/o hardware")

	flag.Parse()

	log.Printf("parsed arguments: %+v", global)
}

func main() {
	parseArguments()

	if !global.NoCheck {
		CheckSystem()
	}

	LoadConfiguration(global.File)
	pinball.initialize()
	pinball.eventLoop()
	pinball.release()
}
