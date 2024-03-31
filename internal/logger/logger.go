package logger

import (
	"log"
	"os"
)

type Logger struct {}

func CreateNew() *log.Logger{
	return log.New(os.Stdout, "", log.Lshortfile)
}