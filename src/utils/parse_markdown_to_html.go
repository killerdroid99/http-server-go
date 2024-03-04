package utils

import (
	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
)

func MarkdownToHTML(md []byte) []byte {
	unsafeHTML := markdown.ToHTML(md, nil, nil)

	html := bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML)

	return html
}
