<div class="starting">
	<div id="tiles" class="grid">
		{{range $index, $element := .data.GridItems}}
			{{ layoutHtml $element }}
		{{end}}
	</div>
</div>
<footer>
{{ componentHtml .data.NavBar}}
</footer>
