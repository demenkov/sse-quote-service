package ssequote

import (
	"log"
)

func LogInfo(msg interface{}, args ...interface{}) {
	log.Println(append([]interface{}{msg}, args...)...)
}

func LogError(msg interface{}, args ...interface{}) {
	if 0 == len(args) || nil == args[0] {
		return
	}

	log.Println(append([]interface{}{msg}, args...)...)
}
