<style>
	#tiles {
		height:100%;
		width:100%;
		display:grid;
		grid-template-columns: {{.data.GridTemplateColumns}};
		grid-template-rows: {{.data.GridTemplateRows}};
		grid-gap: {{ .data.Padding }}px;
		//border: 1px solid #000;
	}
</style>
<div class="starting">
	<div id="tiles" class="primary">
		{{range $index, $element := .data.GridItems}}
			<div class="name" style="border-radius:5px;border:solid 1px #000;text-align:center;padding:2px;grid-area:{{$element.componentData.templateRowStart}}/{{$element.componentData.templateColumnStart}}/{{$element.componentData.templateRowEnd}}/{{$element.componentData.templateColumnEnd}}">
				{{$index}} {{componentHtml $element}}
			</div>
		{{end}}
	</div>
	<HR COLOR="black"><br>

</div>
