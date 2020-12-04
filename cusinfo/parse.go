package main

import (
	"errors"
	"log"
	"strconv"
	"time"

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

// index == 0
const TypeTime = 0

// parse time
// format: 2020-11-28 23:59:06
func HandleTime(str string) (uint64, error) {
	if str == "" {
		return 0, nil
	}
	str = str + " +0800"
	t, err := time.Parse("2006-01-02 15:04:05 -0700", str)
	if err != nil {
		log.Println("parse time error:", str, err)
		return 0, err
	}
	return uint64(t.Unix()), nil
}

const TypeString = 1

// do nothing
func HandleString(str string) string {
	return str
}

const TypeFloat = 2

func HandleFloat(str string) (float32, error) {
	if str == "" {
		return 0, nil
	}
	fl, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, nil
		// log.Println("parse flo;at error:", str, err)
		// return 0, err
	}
	return float32(fl), nil
}
