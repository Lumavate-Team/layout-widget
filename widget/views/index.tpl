<style>
	#tiles {
		height:100%;
		width:100%;
		display:grid;
		grid-template-columns: {{.data.GridTemplateColumns}};
		grid-template-rows: {{.data.GridTemplateRows}};
		grid-gap: {{ .data.Padding }}px;
	}
</style>
<div class="starting">
	<div id="tiles" class="primary">
		{{range $index, $element := .data.GridItems}}
			{{componentHtml $element}}
		{{end}}
	</div>
</div>
