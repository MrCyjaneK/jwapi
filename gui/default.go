// +build !gui,!nogui

package gui

import "log"

func Start() {
	log.Println("[gui] Hello! It looks like you forget to specify a correct tag.")
	log.Println("[gui] Please build me with `-tags gui`, to enable gui or with `-tags nogui` to disable it")
	log.Println("[gui] I'll exit now. Bye.")
}
