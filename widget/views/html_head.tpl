<link href="{{.CacheKey}}/static/css/styles.css" rel="stylesheet">
{{range $i, $src := .data.DirectIncludes }}
  <script src="{{$src}}" type="text/javascript"></script>
{{end}}
<style>
	{{ safeCss .data.InlineCss}}
</style>
