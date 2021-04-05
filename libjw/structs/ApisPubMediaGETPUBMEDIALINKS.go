package structs

// ApisPubMediaGETPUBMEDIALINKS format output from https://app.jw-cdn.org/apis/pub-media/GETPUBMEDIALINKS?langwritten=E&pub=S-38&fileformat=epub
type ApisPubMediaGETPUBMEDIALINKS struct {
	Status struct {
		ID     string `json:"id"`
		Title  string `json:"title"`
		Status int    `json:"status"`
	} `json:"0"`
	PubName       string   `json:"pubName"`
	ParentPubName string   `json:"parentPubName"`
	Booknum       string   `json:"booknum"` // idk, is null
	Pub           string   `json:"pub"`     // the codename
	Issue         string   `json:"issue"`
	FormattedDate string   `json:"formattedDate"`
	FileFormat    []string `json:"fileFormat"`
	Track         string   `json:"track"`     // idk, is null
	Specialty     string   `json:"Specialty"` // idk, is empty string
	PubImage      struct {
		URL              string `json:"url"`
		ModifiedDatetime string `json:"modifiedDateTime"`
		Checksum         string `json:"checksum"`
	} `json:"pubImage"`
	Languages map[string]struct {
		Name      string `json:"name"`
		Direction string `json:"direction"`
		Locale    string `json:"locale"`
	} `json:"languages"`
	Files map[string]map[string][]struct { // s.Files["E"]["EPUB"]... ("E" is the code for english, EPUB is format)
		Title string `json:"title"`
		File  struct {
			URL              string `json:"url"`
			Stream           string `json:"stream"`
			ModifiedDatetime string `json:"modifiedDatetime"`
			Checksum         string `json:"checksum"`
		} `json:"file"`
		Filesize   int
		TrackImage struct {
			URL              string `json:"url"`
			ModifiedDatetime string `json:"modifiedDatetime"`
			Checksum         string `json:"checksum"`
		} `json:"trackImage"`
		Markers        string `json:"markers"` // idk, is null
		Label          string `json:"label"`
		Track          string `json:"track"`
		HasTrack       bool   `json:"hasTrack"`
		Pub            string `json:"pub"`
		Docid          int    `json:"docid"`
		Booknum        int    `json:"booknum"`
		Mimetype       string `json:"mimetype"`
		Edition        string `json:"edition"`
		EditionDescr   string `json:"editionDescr"`
		Format         string `json:"format"`
		FormatDescr    string `json:"formatDescr"`
		Specialty      string `json:"specialty"`
		SpecialtyDescr string `json:"specialtyDescr"`
		Subtitled      bool   `json:"subtitled"`
		FrameWidth     int    `json:"frameWidth"`
		FrameHeight    int    `json:"frameHeight"`
		FrameRate      int    `json:"frameRate"`
		Duration       int    `json:"duration"`
		BitRate        int    `json:"bitRate"`
	} `json:"files"`
}
