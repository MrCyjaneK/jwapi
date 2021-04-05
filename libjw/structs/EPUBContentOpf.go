package structs

import "encoding/xml"

// EPUBContentOpf - This is supposed to read content.obf into some readable thing
type EPUBContentOpf struct {
	XMLName          xml.Name `xml:"package"`
	Version          string   `xml:"version,attr"`
	Xmlns            string   `xml:"xmlns,attr"`
	UniqueIdentifier string   `xml:"unique-identifier,attr"`
	Metadata         struct {
		XMLName   xml.Name `xml:"metadata"`
		DC        string   `xml:"dc,attr"`
		OPF       string   `xml:"opf,attr"`
		DCTitle   string   `xml:"title"`
		DCCreator struct {
			XMLName xml.Name `xml:"creator"`
			Name    string   `xml:",chardata"`
			ID      string   `xml:"id,attr"`
			XMLLang string   `xml:"lang,attr"`
			Dir     string   `xml:"dir,attr"`
		}
		Meta []struct {
			XMLName  xml.Name `xml:"meta"`
			Value    string   `xml:",chardata"`
			Name     string   `xml:"name,attr"`
			Content  string   `xml:"contant,attr"`
			Refines  string   `xml:"refines,attr"`
			Property string   `xml:"property,attr"`
			Scheme   string   `xml:"scheme,attr"`
			ID       string   `xml:"id,attr"`
		} `xml:"meta"` // There are multiple meta tags
		DCLanguage   string `xml:"language"`
		DCPublisher  string `xml:"publisher"`
		DCIdentifier struct {
			XMLName xml.Name `xml:"identifier"`
			ID      string   `xml:"id,attr"`
			Name    string   `xml:",chardata"`
		}
	} // </metadata>
	Manifest struct {
		XMLName xml.Name `xml:"manifest"`
		Items   []struct {
			XMLName    xml.Name `xml:"item"`
			ID         string   `xml:"id,attr"`
			Properties string   `xml:"properties,attr"`
			Href       string   `xml:"href,attr"`
			MediaType  string   `xml:"media-type,attr"`
		} `xml:"item"`
	} // </manifest>
	Spine struct {
		XMLName                  xml.Name `xml:"spine"`
		Toc                      string   `xml:"toc,attr"`
		PageProgressionDirection string   `xml:"page-progression-direction,attr"`
		ItemRefs                 []struct {
			XMLName xml.Name `xml:"itemref"`
			IDRef   string   `json:"idref,attr"`
			Linear  string   `json:"no"`
		} `xml:"itemref"`
	} // </spine>
	Guide struct {
		XMLName xml.Name `xml:"guide"`
		Type    string   `xml:"type,attr"`
		Href    string   `xml:"href,attr"`
		Title   string   `xml:"title,attr"`
	}
}

// note for self dc:title -> title (wtf golang, why?)
