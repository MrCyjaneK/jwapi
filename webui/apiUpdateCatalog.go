package webui

import (
	"log"
	"net/http"
	"strconv"

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
)

func apiUpdateCatalog(w http.ResponseWriter, req *http.Request) {
	log.Println("[webui][apiUpdateCatalog] Updating catalog....")
	for i := range libjw.Alerts {
		if libjw.Alerts[i].Cause == "[libjw][GetCatalog]" {
			libjw.Alerts[i] = libjw.Alert{}
		}
	}
	i := len(libjw.Alerts)
	libjw.Alerts = append(libjw.Alerts, libjw.Alert{
		Title:       "Updating catalog...",
		Description: "Catalog is now being updated, you can continue to use application",
		Color:       "info",
		Cause:       "[webui][apiUpdateCatalog]",
	})
	libjw.GetCatalog(helpers.GetDataDir()+"/raw/catalog.db", false)
	var callbacks []libjw.AlertCallback
	callbacks = append(callbacks, libjw.AlertCallback{
		Title:    "OK",
		Endpoint: "api/alerts/cancel/" + strconv.Itoa(i),
	})
	libjw.Alerts[i] = libjw.Alert{
		Title:       "Catalog updated!",
		Description: "Publications catalog just got updated!",
		Color:       "success",
		Cause:       "[webui][apiUpdateCatalog]",
		Callbacks:   callbacks,
	}
	log.Println("[webui][apiUpdateCatalog] OK!")
}
