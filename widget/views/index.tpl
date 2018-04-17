<style>
	#tiles {
		height:100%;
		width:100%;
		display:grid;
		grid-template-columns: {{.data.GridTemplateColumns}};
		grid-template-rows: {{.data.GridTemplateRows}};
		grid-gap: {{ .data.Padding }}px;
	}
	.nav-tile {
		position:absolute;
		width:100%;
		height:100%;
	}
	.video-frame {
	height:100%;
	width:100%;
	}
	.fill {
		background-position:center;
		background-size: cover;
	}
	.fit {
		background-position:center;
		background-size: contain;
	}
	.stretch {
		background-size: 100% 100%;
	}
	.repeat {
		background-position:center;
		background-repeat:repeat;
	}
</style>
<script>
	function navigate(url) {
		window.location.href = url;
	}
</script>
<div class="starting">
	<div id="tiles" class="primary">
		{{range $index, $element := .data.GridItems}}
			{{componentHtml $element}}
		{{end}}
	</div>
</div>
