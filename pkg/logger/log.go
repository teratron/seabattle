package logger

import (
	"log"
	"os"
)

type Logger struct {
	Info, Warning, Error, Debug *log.Logger
}

var Info = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

//
var Warning = log.New(os.Stdout, "WARNING\t", log.Ldate|log.Ltime|log.Llongfile)

//
var Error = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

//
var Debug = log.New(os.Stderr, "DEBUG\t", log.Ldate|log.Ltime|log.Llongfile)

func New() *log.Logger {
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { err = f.Close() }()

	_ = log.New(f, "INFO\t", log.Ldate|log.Ltime)
	//
	return nil
}
