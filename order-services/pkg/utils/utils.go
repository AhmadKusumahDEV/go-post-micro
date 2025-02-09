package utils

import "log"

func Err(e error, msg string) {
	if msg == "" {
		log.Fatalln(msg, e)
	} else {
		log.Println(e)
	}
}
