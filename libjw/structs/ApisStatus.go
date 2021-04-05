package structs

// ApisStatus - this hold status of api, when something goes wrong
type ApisStatus struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status int    `json:"status"`
}
