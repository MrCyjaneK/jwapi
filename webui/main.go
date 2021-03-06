package webui

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
	htmldata "git.mrcyjanek.net/mrcyjanek/jwapi/webui/html"
)

var SPort = "0"
var Port = 0

// Start the webui
func Start() {
	if SPort == "0" {
		Port = 2000 + rand.Intn(10000)
	} else {
		Port, _ = strconv.Atoi(SPort)
	}
	http.Handle("/", http.FileServer(http.FS(htmldata.Files)))
	http.HandleFunc("/api/", api)
	http.HandleFunc("/api/alerts", apiAlerts)
	http.HandleFunc("/api/alerts/cancel/", apiAlertsCancel)
	http.HandleFunc("/api/db/get/", apiDBget)
	http.HandleFunc("/api/db/set/", apiDBset)
	http.HandleFunc("/api/getIP", apiGetIP)
	http.HandleFunc("/api/publicationList/", apiPublicationList)
	http.HandleFunc("/api/publications/", apiPublications)
	http.HandleFunc("/api/publications_index/", apiPublicationIndex)
	http.HandleFunc("/api/publications_index_toc/", apiPublicationIndexToc)
	http.HandleFunc("/api/publications_json/", apiPublicationJson)
	http.HandleFunc("/api/ping", apiPing)
	http.HandleFunc("/api/updateCatalog", apiUpdateCatalog)
	http.HandleFunc("/api/languages", apiLanguages)
	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(Port), nil)
		if err != nil {
			log.Fatal("[webui][Start]", err)
		}
	}()
	fmt.Println("[webui][Start] Listening on 127.0.0.1:" + strconv.Itoa(Port))
}

func apiPing(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "pong")
}

func api(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "ok | <a href=\"/\">I'll take you home</a>")
}

func apiLanguages(w http.ResponseWriter, req *http.Request) {
	res, err := json.Marshal(libjw.GetLanguages())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(res))
}

/*
func getCookie(req *http.Request, name string) string {
	if len(req.Header["Cookie"]) > 0 {
		cookies := strings.Split(req.Header["Cookie"][0], "; ")
		for i := range cookies {
			cookie := strings.Split(cookies[i], "=")
			if cookie[0] == name {
				return strings.Join(append(cookie[:0], cookie[1:]...), "=")
			}
		}
	}
	return ""
}
*/
