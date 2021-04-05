package structs

// DBPublicationRootKey -
type DBPublicationRootKey struct {
	Symbol   string `json:"symbol"`
	Year     int    `json:"year"` // notnull 0 | should be smallint (int8)
	Language int    `json:"language"`
	ID       int    `json:"id"`
}
