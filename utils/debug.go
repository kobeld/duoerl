package utils

import (
	"log"
	"runtime/debug"
)

func PrintStackAndError(err error) {
	log.Printf("********** Debug Error message: %+v ***********\n", err)
	debug.PrintStack()
}
