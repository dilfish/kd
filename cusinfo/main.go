package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/dilfish/tools"
)

var flagExcelFile = flag.String("f", "", "execl file name")
var flagSrvPort = flag.Int("p", 0, "server port")
var flagConfig = flag.String("c", "", "config file name")
var flagTableName = flag.String("t", "", "table name")

func HandleFlag(fn string) (tools.DBConfig, error) {
	var conf tools.DBConfig
	file, err := os.Open(fn)
	if err != nil {
		log.Println("open file error:", fn, err)
		return conf, err
	}
	defer file.Close()
	bt, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("read all error:", fn, err)
		return conf, err
	}
	err = json.Unmarshal(bt, &conf)
	if err != nil {
		log.Println("unjson error:", fn, err)
		return conf, err
	}
	conf.Ext = "charset=utf8mb4&collation=utf8mb4_unicode_ci"
	return conf, nil
}

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conf, err := HandleFlag(*flagConfig)
	if err != nil {
		log.Println("配置文件错误：", err)
		return
	}
	if *flagExcelFile != "" {
		s := NewService(&conf, *flagTableName)
		if s == nil {
			log.Println("connect db error:", conf)
			return
		}
		err := s.Do(*flagExcelFile)
		log.Println("insert result:", err)
		return
	}
	if *flagSrvPort != 0 {
		s := NewService(&conf, *flagTableName)
		if s == nil {
			log.Println("connect db error:", conf)
			return
		}
		err := s.Srv(*flagSrvPort)
		log.Println("server mode:", err)
	}
	return
}
