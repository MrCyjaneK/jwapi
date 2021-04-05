package jwhttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// - CatalogsPublicationsV4Manifest is meant to receive part of url required to get catalog.db
// - SiteLanguages fetch https://www.jw.org/en/languages/ and return structs.Languages{}
// - CatalogsPublicationsV4CatalogDbGz get gzipped catalog'https://app.jw-cdn.org/catalogs/publications/v4/815ecc5a-e72b-48ac-aa9f-e1e7d256e43b/catalog.db.gz'
// - ApisPubMediaGETPUBMEDIALINKS - Fetch publication links

//curl --silent 'https://app.jw-cdn.org/catalogs/publications/v4/815ecc5a-e72b-48ac-aa9f-e1e7d256e43b/catalog.info.json.gz' | zcat
//{
//  "size": 32130196,
//  "revision": 1868639,
//  "schema": 4,
//  "created": "2021-01-04T05:39:02+00:00"
//}⏎

func httpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}
