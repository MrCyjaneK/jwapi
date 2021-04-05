package structs

// CatalogsPublicationsV4Manifest is meant to receive part of url required to get catalog.db
//curl --silent https://app.jw-cdn.org/catalogs/publications/v4/manifest.json | jq
//{
//	"version": 1,
//	"current": "815ecc5a-e72b-48ac-aa9f-e1e7d256e43b"
//}
type CatalogsPublicationsV4Manifest struct {
	Version int    `json:"version"`
	Current string `json:"current"`
}
