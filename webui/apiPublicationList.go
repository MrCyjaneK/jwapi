package webui

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw/structs"
)

type Publication struct {
	Title string `json:"title"`
	Code  string `json:"code"`
}

func apiPublicationList(w http.ResponseWriter, req *http.Request) {
	jsonData, err := ioutil.ReadFile(helpers.GetDataDir() + "/catalog/Publication.json")
	if err != nil {
		log.Fatal(err)
	}

	var publications []structs.DBPublication
	err = json.Unmarshal(jsonData, &publications)
	if err != nil {
		log.Fatal(err)
	}
	var mylang int
	for j := range libjw.MepsMap {
		if libjw.MepsMap[j] == string(helpers.Get("lang")) {
			mylang = j
			break
		}
	}
	var pubs []Publication
	for i := range publications {
		p := publications[i]
		title := "undefined"
		code := p.KeySymbol
		if p.MepsLanguageID != mylang {
			continue
		}
		if p.Title != "" {
			title = p.Title
		}
		if p.IssueTitle != "" {
			title = p.IssueTitle
		}
		if p.IssueTagNumber != 0 {
			code = code + "_" + strconv.Itoa(p.IssueTagNumber)
		}

		if title != "undefined" && code != "undefined" {
			pubs = append(pubs, Publication{
				Title: title,
				Code:  code,
			})
		}
	}
	resp, err := json.Marshal(pubs)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}
