package htmldata

import "embed"

//go:embed *.html css img static vendor
var Files embed.FS
