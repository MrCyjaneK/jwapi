package structs

// DBPublicationAsset -
type DBPublicationAsset struct {
	PublicationID              int    `json:"publication_id"`
	MepsLanguageID             int    `json:"meps_language_id"`
	Signature                  string `json:"signature"`
	Size                       int    `json:"size"`
	ExpandedSize               int    `json:"expanded_size"`
	MimeType                   string `json:"mime_type"`
	SchemaVersion              int    `json:"schema_version"`
	MinPlatformVersion         int    `json:"min_platform_version"`
	CatalogedOn                string `json:"cataloged_on"`
	LastUpdated                string `json:"last_updated"`
	LastModified               string `json:"last_modified"`
	GenerallyAvailableDate     string `json:"generally_available_date"`      // notnull 0
	ConventionReleaseDayNumber int    `json:"convention_release_day_number"` // notnull 0
	ID                         int    `json:"id"`
}
