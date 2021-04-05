package main

import (
	"time"

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
	"git.mrcyjanek.net/mrcyjanek/jwapi/webui"
)

var (
	dataDir = ""
)

func main() {
	if dataDir == "" {
		dataDir = helpers.GetDataDir()
	} else {
		helpers.SetDataDir(dataDir)
	}
	//go func() {
	//	for {
	//		helpers.PrintMemUsage()
	//		time.Sleep(time.Millisecond * 500)
	//	}
	//}()
	helpers.Mkdir(dataDir + "/raw")
	helpers.DBInit(dataDir)
	libjw.GetCatalog(dataDir + "/raw/catalog.db")
	webui.Start()
	//w := webview.New(true) // false - debug
	//defer w.Destroy()
	//w.SetTitle("openJW Library")
	//w.SetSize(800, 600, webview.Hint(webview.HintNone))
	//w.Navigate("http://127.0.0.1:8080")
	//w.Run()
	for {
		time.Sleep(time.Hour)
	}
}
