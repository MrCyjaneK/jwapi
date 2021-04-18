package gui

import "log"

func Start() {
	log.Println("Hello! It looks like you forget to specify a correct tag.")
	log.Println("Please build me with `-tags gui`, to enable gui or with `-tags nogui` to disable it")
	log.Println("I'll exit now. Bye.")
}
