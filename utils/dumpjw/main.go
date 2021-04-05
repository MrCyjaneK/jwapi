package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw/structs"
)

func main() {
	var lang = "E"

	if lang == "" {
		fmt.Println("Lang not provided", lang)
		os.Exit(1)
	}

	var publicationRootKey []structs.DBPublicationRootKey
	var bytes, err = ioutil.ReadFile(helpers.GetDataDir() + "/catalog/PublicationRootKey.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	json.Unmarshal(bytes, &publicationRootKey)
	for i := range publicationRootKey {
		pub := publicationRootKey[i]
		libjw.GetPublication(pub.Symbol, lang, "EPUB", "")
	}
	//{
	//	"symbol": "iagwa",
	//	"year": 0,
	//	"language": 0,
	//	"id": 0
	//},
}
