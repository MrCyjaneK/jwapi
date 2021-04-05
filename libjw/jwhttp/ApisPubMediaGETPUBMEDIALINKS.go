package jwhttp

import (
	"encoding/json"
	"fmt"

	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw/structs"
)

// ApisPubMediaGETPUBMEDIALINKS - Fetch publication links
func ApisPubMediaGETPUBMEDIALINKS(langwritten string, pub string, fileformat string, issue string) (toret structs.ApisPubMediaGETPUBMEDIALINKS, status []structs.ApisStatus) {
	toret = structs.ApisPubMediaGETPUBMEDIALINKS{}
	url := ""
	if issue == "" {
		url = "https://app.jw-cdn.org/apis/pub-media/GETPUBMEDIALINKS?langwritten=" + langwritten + "&pub=" + pub + "&fileformat=" + fileformat
	} else {
		url = "https://app.jw-cdn.org/apis/pub-media/GETPUBMEDIALINKS?langwritten=" + langwritten + "&pub=" + pub + "&fileformat=" + fileformat + "&issue=" + issue
	}
	fmt.Println(url)
	resp := httpGet(url)
	json.Unmarshal(resp, &toret)
	status = []structs.ApisStatus{}
	json.Unmarshal(resp, &status) // same shit, different story
	return toret, status
}
