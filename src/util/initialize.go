package util

import (
	"flag"
	"os"

	"github.com/Unknwon/goconfig"
	log "github.com/cihub/seelog"
)

var ExitChan chan int
var IsDebug bool

var confFile string

func init() {
	ExitChan = make(chan int)

	flag.StringVar(&confFile, "conf", "./src/config/config.ini", "-conf=/path/to/config.ini")
	flag.BoolVar(&IsDebug, "debug", false, "-debug")
	flag.Parse()
}

func LoadConfig() map[string]string {
	if _, err := os.Stat(confFile); err != nil {
		log.Criticalf("config file[%s] not exist: %s", confFile, err)
		os.Exit(2)
	}

	configFile, err := goconfig.LoadConfigFile(confFile)
	if err != nil {
		log.Criticalf("load config file err: %s", err)
		os.Exit(2)
	}
	fieldList := configFile.GetKeyList("")

	conf := make(map[string]string)
	for _, field := range fieldList {
		if conf[field], err = configFile.GetValue("", field); err != nil {
			log.Criticalf("parse field %s error: %s", field, err)
			os.Exit(2)
		}
	}
	return conf
}
