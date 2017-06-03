package micha

import (
	"log"
	"os"
)

// Logger interface
type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

func newLogger(prefix string) Logger {
	return log.New(os.Stderr, prefix, log.Ldate|log.Ltime|log.Lshortfile)
}
