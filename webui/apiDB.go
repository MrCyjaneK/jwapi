package webui

import (
	"fmt"
	"net/http"
	"strings"

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
)

func apiDBget(w http.ResponseWriter, req *http.Request) {
	w.Write(helpers.Get(req.URL.RawQuery))
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
	helpers.Set(key, []byte(query))

}
