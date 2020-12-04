package main

import (
	"errors"
	"log"
	"github.com/tealeg/xlsx"
)

func (s *Service) Do(fn string) error {
	slice, err := xlsx.FileToSlice(fn)
	if err != nil {
		log.Println("excel to slice error:", err)
		return err
	}
	if len(slice) != 1 {
		log.Println("we have some inner error:", len(slice))
		return errors.New("slice is not len 1")
	}
	table := slice[0]
	log.Println("we have", len(table), "lines of data")
	for idx, t := range table {
		// jump table header
		if idx == 0 {
			continue
		}
		if (idx % 50) == 0 {
			log.Println("we have inserted", idx, "records")
		}
		err := s.InsertDB(t)
		if err != nil {
			log.Println("insert db error:", idx, err)
			return err
		}
	}
	return nil
}
