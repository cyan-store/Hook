package log

import (
	"log"
	"os"
)

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

func init() {
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate)
	Info = log.New(os.Stdout, "Info: ", log.Ldate)
	Error = log.New(os.Stdout, "Error: ", log.Ldate)
}
