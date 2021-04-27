// +build guilorca

package gui

import (
	"strconv"

	"git.mrcyjanek.net/mrcyjanek/jwapi/webui"
	"github.com/zserge/lorca"
)

func Start() {
	ui, _ := lorca.New("", "", 480, 320)
	ui.Eval(`window.location.href = "http://127.0.0.1:` + strconv.Itoa(webui.Port) + `"`)
	<-ui.Done()
}
