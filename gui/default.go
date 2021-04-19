// +build !gui,!nogui

package gui

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"git.mrcyjanek.net/mrcyjanek/jwapi/webui"
	"github.com/pkg/browser"
)

func Start() {
	log.Println("[gui] Trying to start frontend...")
	err := browser.OpenURL("http://127.0.0.1:" + strconv.Itoa(webui.Port))
	if err != nil {
		log.Println("[gui]", err)
	}
	log.Println("[gui] Please stop me after closing app.")
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-cancelChan
	log.Printf("[gui] Caught SIGTERM %v, exiting", sig)
}
