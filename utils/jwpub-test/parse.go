package main

import (
	"log"
	"os"

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
)

func main() {
	dataDir := helpers.GetDataDir()
	helpers.SetDataDir(dataDir)
	helpers.Mkdir(dataDir + "/raw")
	helpers.DBInit(dataDir)
	//libjw.GetPublication("w", "E", "JWPUB", "202110")
	if _, err := os.Stat("pub.jwpub"); os.IsNotExist(err) {
		log.Fatal("Hey! Please put `pub.jwpub' in this directory, you can get one from this link: https://www.jw.org/download/?issue=202107&output=html&pub=g&fileformat=JWPUB&alllangs=0&langwritten=E&txtCMSLang=E&isBible=0")
	}
	libjw.JWPUBtoMarkdown("pub.jwpub")
}
