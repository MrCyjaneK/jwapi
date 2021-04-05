package structs

// PublicationV2 stores data only for specific publication (nwt_P instead of nwt)
type PublicationV2 struct {
	Title  string `json:"title"`
	Format string `json:"format"`
	Path   string `json:"path"` // keep part of path.
	URI    string `json:"uri"`
}
