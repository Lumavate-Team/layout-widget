<link href="{{.CacheKey}}/static/css/styles.css" rel="stylesheet">
<style>
	.grid {
		position:absolute;
		height:100%;
		width:100%;
		display:grid;
		grid-template-columns: {{.data.GridTemplateColumns}};
		grid-template-rows: {{.data.GridTemplateRows}};
		grid-gap: {{.data.Padding}}px;
	}
	@supports not (grid-area: 1/1/1/1) {
		.grid-item {
			margin: {{.data.Padding}}px;
		}
	}
	body {
		position:absolute;
		height: 100%;
		width:100%;
		{{ if .data.DisplayBackgroundImage }}
		background-image: url({{.data.BackgroundImage.Preview}});
		background-repeat:repeat;
		{{ else if .data.BackgroundColor }}
		background-color: {{.data.BackgroundColor}};
		{{ end }}
	}
</style>
