package structs

// DBPublication -
type DBPublication struct {
	PublicationRootKeyID  int    `json:"publication_root_key_id"`
	MepsLanguageID        int    `json:"meps_language_id"`
	PublicationTypeID     int    `json:"publication_type_id"`
	IssueTagNumber        int    `json:"issue_tag_number"`
	Title                 string `json:"title"`
	IssueTitle            string `json:"issue_title"` // notnull 0
	ShortTitle            string `json:"short_title"`
	CoverTitle            string `json:"cover_title"`             // notnull 0
	UndatedTitle          string `json:"undated_title"`           // notnull 0
	UndatedReferenceTitle string `json:"undated_reference_title"` // notnull 0
	Year                  int    `json:"year"`                    // should be smallint (int8), I'll go for int now
	Symbol                string `json:"symbol"`
	KeySymbol             string `json:"key_symbol"` // notnull 0 | varchar(16)
	Reserved              int    `json:"reserved"`
	ID                    int    `json:"id"`
}
