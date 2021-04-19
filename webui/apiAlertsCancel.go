package webui

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
)

func apiAlertsCancel(w http.ResponseWriter, req *http.Request) {
	splited := strings.Split(string(req.URL.Path), "/")
	if len(splited) == 5 {
		num, err := strconv.ParseInt(splited[4], 10, 32)
		if err != nil {
			log.Println("[webui][apiAlertsCancel]", err)
			return
		}
		if int64(len(libjw.Alerts)) >= num {
			libjw.Alerts[num] = libjw.Alert{}
		}
	}
}
