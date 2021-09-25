package log

import (
	"fmt"
	"log"
	"os"
)

var Log *log.Logger

func init() {
	fmt.Println("Creating New Instance of Logger....")
	NewLogger()
}

// Creating new logger
func NewLogger() {
	Log = log.New(os.Stderr, "baston-eventos", 1)
}
