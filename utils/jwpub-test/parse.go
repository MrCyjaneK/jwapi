package main

import (
	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
)

func main() {
	dataDir := helpers.GetDataDir()
	helpers.SetDataDir(dataDir)
	helpers.Mkdir(dataDir + "/raw")
	helpers.DBInit(dataDir)
	libjw.GetPublication("w", "E", "JWPUB", "202110")
	// libjw.JWPUBtoMarkdown("fg_E.jwpub.orig")
}
