package structs

// Publication holds all kind of information about a publication
type Publication struct {
	Code     string // This is meant to store the codename of publication (for example 'nwt' is the Bible)
	Year     int
	Image    Img
	LangData *[]LangData
}
