package webui

import (
	"encoding/json"
	"log"
	"net/http"

	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
)

func apiAlerts(w http.ResponseWriter, req *http.Request) {
	resp, err := json.Marshal(libjw.Alerts)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}
