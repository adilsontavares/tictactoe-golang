package client

import (
	"log"
	"os"
)

var (
	file *os.File
)

func InitLog(path string) bool {

	var err error
	file, err = os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		return false
	}

	log.SetOutput(file)
	log.Println("Log initialized.")

	return true
}

func CloseLog() {

	if file != nil {
		
		file.Close()
		file = nil
	}
}