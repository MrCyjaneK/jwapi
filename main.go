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
	helpers.Mkdir(dataDir + "/raw")
	helpers.DBInit(dataDir)
	webui.Start()
	go libjw.GetCatalog(dataDir+"/raw/catalog.db", true)
	gui.Start()
}
