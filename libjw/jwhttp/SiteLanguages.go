package jwhttp

import (
	"encoding/json"

	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw/structs"
)

// SiteLanguages - fetch https://www.jw.org/en/languages/ and return structs.Languages{}
func SiteLanguages(baselang string) structs.Languages {
	toret := structs.Languages{}
	json.Unmarshal(httpGet("https://www.jw.org/"+baselang+"/languages/"), &toret)
	return toret
}
