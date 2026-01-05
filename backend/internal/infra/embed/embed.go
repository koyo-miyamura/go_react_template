package embed

import "embed"

//go:embed all:dist
var DistFS embed.FS

//go:embed swagger/swagger-ui.html
var SwaggerUIHTML []byte
