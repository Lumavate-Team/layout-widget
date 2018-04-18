<script>
	function navigate(url) {
		window.location.href = url;
	}
</script>
<div class="starting">
	<div id="tiles" class="primary">
		{{range $index, $element := .data.GridItems}}
			{{ layoutHtml $element}}
		{{end}}
	</div>
</div>
<footer>
{{ componentHtml .data.NavBar}}
</footer>
