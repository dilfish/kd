package main

import (
	"flag"
	"log"

	"github.com/dilfish/tools"
)

var flagExcelFile = flag.String("f", "", "execl file name")
var flagSrvMode = flag.Bool("s", false, "server mode")
var flagSrvPort = flag.Int("port", 0, "server port")
var flagDB = flag.String("d", "", "dbname")
var flagHost = flag.String("h", "", "host ip")
var flagPort = flag.Int("p", 0, "port number")
var flagPass = flag.String("P", "", "password")
var flagUser = flag.String("u", "", "user name")

func HandleFlag() tools.DBConfig {
	var conf tools.DBConfig
	conf.DBName = *flagDB
	conf.Host = *flagHost
	conf.Port = *flagPort
	conf.User = *flagUser
	conf.Pass = *flagPass
	// conf.DBName = "kd"
	// conf.Host = "127.0.0.1"
	// conf.Pass = ""
	// conf.Port = 3306
	// conf.User = "root"
	conf.Ext = "charset=utf8mb4&collation=utf8mb4_unicode_ci"
	return conf
}

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conf := HandleFlag()
	if *flagExcelFile != "" {
		s := NewService(&conf)
		if s == nil {
			log.Println("connect db error:", conf)
			return
		}
		err := s.Do(*flagExcelFile)
		log.Println("insert result:", err)
		return
	}
	if *flagSrvMode {
		s := NewService(&conf)
		if s == nil {
			log.Println("connect db error:", conf)
			return
		}
		err := s.Srv(*flagSrvPort)
		log.Println("server mode:", err)
	}
	return
}
