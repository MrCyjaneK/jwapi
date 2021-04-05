package structs

// LangData stores per-language data
type LangData struct {
	Code          string // E - English | P - Polish
	Title         string
	InitialFormat string // the initial format, like epub
	Path          string // This should contain part of path ~/LibJWgo/publications should be skipped, and only /nwt/P/ should be given
}
