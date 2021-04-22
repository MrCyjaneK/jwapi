// +build gui-webview

package gui

import (
	"strconv"

	"git.mrcyjanek.net/mrcyjanek/jwapi/webui"
	"github.com/webview/webview"
)

func Start() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("JW Study")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://127.0.0.1:" + strconv.Itoa(webui.Port))
	w.Run()
}
