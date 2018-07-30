{{range $index, $element := .data.BodyItems}}
  {{ layoutHtml $element }}
{{end}}

