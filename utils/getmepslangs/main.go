package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw/structs"
)

var (
	dataDir = ""
)

func main() {
	if dataDir == "" {
		helpers.SetDataDir(filepath.Join(os.TempDir(), "/LibJWgo"))
	}

	dataDir = helpers.GetDataDir()

	helpers.Mkdir(dataDir + "/raw")
	helpers.DBInit(dataDir)
	libjw.GetCatalog(dataDir+"/raw/catalog.db", false)
	var publications []structs.DBPublication

	jsonData, err := ioutil.ReadFile(dataDir + "/catalog/Publication.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonData, &publications)
	if err != nil {
		log.Fatal(err)
	}
	// Get list of all publications
	/*
		var pubamt = make(map[string]int)
		for i := range publications {
			pub := publications[i]
			if pub.KeySymbol == "" {
				continue
			}
			pubamt[pub.KeySymbol]++
			log.Println(pub.KeySymbol)
		}
		var biggest string
		var biggestamt int
		for j := range pubamt {
			pk := pubamt[j]
			if pk < biggestamt {
				biggest = j
			}
		}
	*/
	// Now fetch this publication in all the languages,
	//                  langcode meps
	var langs = make(map[string]int)
	biggest := "t-31"
	var langlist = libjw.GetLanguages().Languages
	for k := range langlist {
		l := langlist[k]
		langs[l.Langcode] = -1
	}
	os.Stderr.WriteString(`package libjw
// This file is autogenerated. Please do not submit pull requests to it.
// If you found that some Meps language ID is not equal to actual language
// please edit /utils/getmepslangs/main.go
// Look for comment saying 'Hardcode languages here'
// 
// This file was generated on ` + time.Now().String() + `

var MepsMap = map[int]string{
`)
	// Hardcode languages here
	os.Stderr.WriteString("\t198: \"P\",\t// HARDCODED: 198 -> Polish\n")
	langs["P"] = 198
	// Simply duplicate these 2 lines and replace what's needed
	os.Stderr.WriteString("\t0: \"E\",\t// HARDCODED: 0 -> English\n")
	langs["E"] = 0

	for k := range langlist {
		l := langlist[k]
		if langs[l.Langcode] != -1 {
			continue
		}
		pub, err := libjw.GetPublication(biggest, l.Langcode, "EPUB", "")
		if err != nil {
			continue
		}
		for i := range publications {
			pub2 := publications[i]
			leng := 16
			if len(pub.Title) < leng {
				leng = len(pub.Title) - 5
			}
			if len(pub2.Title) < leng {
				leng = len(pub2.Title) - 5
			}
			if pub.Title[0:leng] == pub2.Title[0:leng] &&
				helpers.StrCompare(pub.Title, pub2.Title) > 50 {
				//log.Println(pub.Title, "==", pub2.Title)
				//log.Println(pub2.MepsLanguageID, "==", l.Langcode, "(score:", helpers.StrCompare(pub.Title, pub2.Title), ")")
				b := false
				for j := range langs {
					if langs[j] == pub2.MepsLanguageID {
						b = true
					}
				}
				if b {
					continue
				}
				langs[l.Langcode] = pub2.MepsLanguageID
				os.Stderr.WriteString("\t" + strconv.Itoa(pub2.MepsLanguageID) + ": " + "\"" + l.Langcode + "\",\t// AUTOGEN: '" + pub.Title + "' == '" + pub2.Title + "' (score:" + strconv.FormatFloat(helpers.StrCompare(pub.Title, pub2.Title), 'f', 2, 64) + ")\n")
				break
			}
		}
	}
	os.Stderr.WriteString("}")
}
