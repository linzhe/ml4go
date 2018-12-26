package dataset

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger = log.New(os.Stdout, "", 0)

func SetLogger(logger *log.Logger) {
	Logger = logger
}

func SetLoggerOut(out io.Writer) {
	Logger = log.New(out, "", 0)
}

func Silent() {
	if out, err := os.Open(os.DevNull); err != nil {
		panic(err)
	} else {
		Logger = log.New(out, "", 0)
	}
}
