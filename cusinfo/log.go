package main

import (
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func NewLog(addr string) *log.Logger {
	prefix := time.Now().Format("2006-01-02 15:04:05 " + addr + "-")
	uu, err := uuid.NewRandom()
	if err == nil {
		prefix = uu.String() + "-"
	}
	xl := log.New(os.Stderr, prefix, log.LstdFlags|log.Lshortfile)
	xl.Println("新的日志对象:", addr, prefix)
	return xl
}
