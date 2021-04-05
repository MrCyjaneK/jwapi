package structs

// DBImageAsset -
type DBImageAsset struct {
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	NameFragment string `json:"name_fragment"`
	Size         int    `json:"size"`
	MimeType     string `json:"mime_type"`
	ID           int    `json:"id"`
}
