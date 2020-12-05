package main

import (
	"errors"
	"log"

	"github.com/tealeg/xlsx"
)

var ErrBadName = errors.New("excel 文件名字错误")

func (s *Service) Do(fn string) error {
	err := s.CreateTable()
	if err != nil {
		log.Println("创建数据库表错误:", err)
		return err
	}
	slice, err := xlsx.FileToSlice(fn)
	if err != nil {
		log.Println("打开 excel 文件错误:", err)
		return err
	}
	if len(slice) != 1 {
		log.Println("excel 文件格式错误:", len(slice))
		return errors.New("excel 文件格式错误")
	}
	table := slice[0]
	log.Println(fn, "文件有", len(table), "行数据")
	for idx, t := range table {
		// jump table header
		if idx == 0 {
			continue
		}
		if (idx % 50) == 0 {
			log.Println("已经写入", idx, "条")
		}
		err := s.InsertDB(t)
		if err != nil {
			log.Println("写入数据库错误:", idx, err)
			return err
		}
	}
	return nil
}
