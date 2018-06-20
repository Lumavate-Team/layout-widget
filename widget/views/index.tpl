{{range $index, $element := .data.GridItems}}
  {{ layoutHtml $element }}
{{end}}

