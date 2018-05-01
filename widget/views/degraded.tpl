<script>
	function navigate(url) {
		window.location.href = url;
	}
</script>
<style>
	.grid {
		display:flex;
		flex-direction:column;
	}
	.optimal {
		display:none;
	}
	.degraded {
		display: flex;
		flex:1;
		vertical-align: top;
	}
</style>
<div class="starting">
	<div id="tiles" class="grid">
		{{range $index, $element := .data.GridItems}}
			{{ if ne $element.DisplayMode "optimal" }}
				{{ layoutHtml $element}}
			{{ end }}
		{{end}}
	</div>
</div>
<footer>
{{ componentHtml .data.NavBar}}
</footer>
