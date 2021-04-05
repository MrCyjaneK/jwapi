package structs

// DBDatedText -
type DBDatedText struct {
	Class         int    `json:"class"`
	Start         string `json:"start"`
	End           string `json:"end"`
	PublicationID int    `json:"publication_id"`
}
