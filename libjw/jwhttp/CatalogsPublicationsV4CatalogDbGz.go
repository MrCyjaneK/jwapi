package jwhttp

//CatalogsPublicationsV4CatalogDbGz get gzipped catalog
// 'https://app.jw-cdn.org/catalogs/publications/v4/815ecc5a-e72b-48ac-aa9f-e1e7d256e43b/catalog.db.gz'
func CatalogsPublicationsV4CatalogDbGz(version string) []byte {
	return httpGet("https://app.jw-cdn.org/catalogs/publications/v4/" + version + "/catalog.db.gz")
}
