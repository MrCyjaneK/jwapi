package webui

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
)

func apiDBget(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	splited := strings.Split(string(url), "/")
	if len(splited) < 5 {
		fmt.Fprintln(w, "/api/db/get/<key>")
		return
	}
	key := splited[4]
	if key == "" {
		w.Write([]byte("0"))
		return
	}
	w.Write(helpers.Get(key))
}

func apiDBset(w http.ResponseWriter, req *http.Request) {
	urla := req.URL.Path
	splited := strings.Split(string(urla), "/")
	if len(splited) < 5 {
		fmt.Fprintln(w, "/api/db/set/<key>?value")
		return
	}
	key := splited[4]
	if key == "" {
		w.Write([]byte(""))
		return
	}
	query := req.URL.RawQuery
	value, err := url.QueryUnescape(query)
	if err != nil {
		log.Fatal(err)
	}
	helpers.Set(key, []byte(value))

}
