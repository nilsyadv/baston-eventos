package log

import (
	"log"
	"os"
)

var Log *log.Logger

func init() {
	NewLogger()
	Log.Println("New Instance Of Logger Created Successfully.")
}

// Creating new logger
func NewLogger() {
	Log = log.New(os.Stderr, "baston-eventos ", 1)
}
