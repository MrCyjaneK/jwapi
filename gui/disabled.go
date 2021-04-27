// +build nogui

package gui

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	log.Println("[gui] Compile with `-tags gui{browser,lorca,webview}` to enable gui, waiting on a infinite loop.")
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-cancelChan
	log.Printf("[gui] Caught SIGTERM %v, exiting", sig)
}
