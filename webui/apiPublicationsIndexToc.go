package webui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw/structs"
)

// This one is responsble for displaying the index of publications
// including things like all chapters and information if a chapter
// should be displayed in navigation

func apiPublicationIndexToc(w http.ResponseWriter, req *http.Request) {
	lang := string(helpers.Get("lang"))
	if lang == "" {
		return
	}
	url := req.URL.Path
	splited := strings.Split(string(url), "/")
	publication := splited[3]
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		fmt.Println(err)
		return
	}
	//extractpath := datadir + "/data/publications/" + publication + "/"
	//helpers.Mkdir(extractpath)
	issue := ""
	pubExploded := strings.Split(publication, "_")
	if len(pubExploded) != 1 {
		issue = pubExploded[1]
		publication = pubExploded[0]
	}
	publication = reg.ReplaceAllString(publication, "")
	issue = reg.ReplaceAllString(issue, "")
	pubdata, err := libjw.GetPublication(publication, lang, "EPUB", issue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "<b>bad request</a> <a href=\"..\">go back</a>")
		return
	}
	contentTocPath := helpers.GetDataDir() + pubdata.Path + "OEBPS/toc.ncx"
	//fmt.Fprintln(w, contentOpfPath)

	content := libjw.DecodeContentToc(contentTocPath)
	//fmt.Println(len(splited))
	var listOfContent []structs.PublicationChapter
	for i := range content.NavMap.NavPoint {
		elm := content.NavMap.NavPoint[i]
		listOfContent = append(listOfContent, structs.PublicationChapter{
			Title: elm.NavLabel.Text,
			URL:   "/api/publications/" + publication + "/" + elm.Content.Src,
		})
	}
	jbytes, err := json.Marshal(listOfContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("[webui][apiPublicationIndex] failed to json.Marshal(listOfContent)")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Cache-Control", "public")
	io.Copy(w, bytes.NewReader(jbytes))
}
