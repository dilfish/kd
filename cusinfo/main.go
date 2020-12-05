package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/dilfish/tools"
)

var flagExcelFile = flag.String("f", "", "execl 文件名")
var flagSrvPort = flag.Int("p", 0, "监听端口")
var flagConfig = flag.String("c", "./config.conf", "config file name")
var flagTableName = flag.String("t", "", "数据库表名，推荐和 excel 文件一样，但不要后缀的 xlsx")

func HandleFlag(fn string) (tools.DBConfig, error) {
	var conf tools.DBConfig
	file, err := os.Open(fn)
	if err != nil {
		log.Println("打开文件错误:", fn, err)
		return conf, err
	}
	defer file.Close()
	bt, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("读取文件错误:", fn, err)
		return conf, err
	}
	err = json.Unmarshal(bt, &conf)
	if err != nil {
		log.Println("解压 json 错误:", fn, err)
		return conf, err
	}
	conf.Ext = "charset=utf8mb4&collation=utf8mb4_unicode_ci"
	return conf, nil
}

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := *flagConfig
	if config == "" {
		config = "./config.conf"
	}
	conf, err := HandleFlag(config)
	if err != nil {
		log.Println("配置文件错误：", err)
		return
	}
	if *flagExcelFile != "" {
		s := NewService(&conf, *flagTableName)
		if s == nil {
			log.Println("连接数据库错误:", conf)
			return
		}
		err := s.Do(*flagExcelFile)
		log.Println("写入数据库错误:", err)
		return
	}
	if *flagSrvPort != 0 {
		s := NewService(&conf, *flagTableName)
		if s == nil {
			log.Println("连接数据库错误:", conf)
			return
		}
		err := s.Srv(*flagSrvPort)
		log.Println("服务器启动错误:", err)
	}
	return
}
