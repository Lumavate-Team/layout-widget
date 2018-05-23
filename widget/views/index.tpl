<div class="layout-starting">
	<div id="tiles" class="layout-grid">
		{{range $index, $element := .data.GridItems}}
			{{ layoutHtml $element }}
		{{end}}
	</div>
</div>
<footer>
{{ componentHtml .data.NavBar}}
</footer>
