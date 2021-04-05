package jwhttp

import (
	"encoding/json"

	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw/structs"
)

// CatalogsPublicationsV4Manifest is meant to receive part of url required to get catalog.db
//curl --silent https://app.jw-cdn.org/catalogs/publications/v4/manifest.json | jq
//{
//	"version": 1,
//	"current": "815ecc5a-e72b-48ac-aa9f-e1e7d256e43b"
//}
func CatalogsPublicationsV4Manifest() structs.CatalogsPublicationsV4Manifest {
	toret := structs.CatalogsPublicationsV4Manifest{}
	json.Unmarshal(httpGet("https://app.jw-cdn.org/catalogs/publications/v4/manifest.json"), &toret)
	return toret
}
