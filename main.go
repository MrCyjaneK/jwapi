package main

import (
	"git.mrcyjanek.net/mrcyjanek/jwapi/gui"
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
	gui.Start()
}
