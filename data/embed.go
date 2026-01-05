package data

import "embed"

//go:embed *.json
var EmbeddedFS embed.FS
