package structs

// Languages response from site https://www.jw.org/en/languages/
type Languages struct {
	Status    int `json:"status"` // 200 - OK
	Languages []struct {
		Symbol         string   `json:"symbol"`         // pl
		Langcode       string   `json:"langcode"`       // P <--- this is what you need (nwt_P)
		Name           string   `json:"name"`           // Polish
		VernacularName string   `json:"vernacularName"` // Polish
		AltSpellings   []string `json:"altSpellings"`   // Polish, polski
		Direction      string   `json:"direction"`      // ltr
		IsSignLanguage bool     `json:"isSignLanguage"`
		IsCounted      bool     `json:"IsCounted"`
		HasWebContent  bool     `json:"hasWebContent"`
	} `json:"languages"`
	LocalizedCount string `json:"localizedCount"` // How many languages do we have?
}
