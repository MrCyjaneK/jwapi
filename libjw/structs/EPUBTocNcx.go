package structs

import "encoding/xml"

// EPUBTocNcx - This is supposed to read toc.ncx into some readable thing
type EPUBTocNcx struct {
	XMLName xml.Name `xml:"ncx"`
	Version string   `xml:"version,attr"`
	Lang    string   `xml:"lang,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Head    struct {
		XMLName xml.Name `xml:"head"`
		Meta    []struct {
			XMLName xml.Name `xml:"meta"`
			Name    string   `xml:"name,attr"`
			Content string   `xml:"contant,attr"`
		} `xml:"meta"`
	}
	DocTitle struct {
		XMLName xml.Name `xml:"docTitle"`
		Text    string   `xml:"text"`
	}
	DocAuthor struct {
		XMLName xml.Name `xml:"docAuthor"`
		Text    string   `xml:"text"`
	}
	NavMap struct {
		XMLName  xml.Name `xml:"navMap"`
		NavPoint []struct {
			XMLName   xml.Name `xml:"navPoint"`
			Class     string   `xml:"class,attr"`
			ID        string   `xml:"id,attr"`
			PlayOrder string   `xml:"playOrder,attr"` // Maybe it should be int?
			NavLabel  struct {
				XMLName xml.Name `xml:"navLabel"`
				Text    string   `xml:"text"`
			}
			Content struct {
				XMLName xml.Name `xml:"content"`
				Src     string   `xml:"src,attr"`
			}
		} `xml:"navPoint"`
	}
}

// note for self dc:title -> title (wtf golang, why?)
