<!DOCTYPE html>
<html lang="en">
  <head>
    <title>{{.data.InstanceName}}</title>
    <meta charset="utf-8">
    <base href="{{.WidgetUrlPrefix}}">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <meta name="description" content="{{.data.InstanceName}}">
    <meta name="theme-color" content="#ffffff">
    <noscript>Javascript required for this site to work.</noscript>

    <link rel="apple-touch-icon" sizes="180x180" href="/iot/favicon-180x180.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/iot/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/iot/favicon-16x16.png">
    <link rel="manifest" href="/manifest.json">
    <link rel="shortcut icon" href="/iot/favicon.ico">
    <meta name="apple-mobile-web-app-capable" content="no">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta property="og:image" content="{{.dnsInfo}}/iot/android-chrome-512x512.png" />

    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Teko:400,500" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro" rel="stylesheet">
    <link rel="stylesheet" href="https://unpkg.com/purecss@1.0.0/build/pure-min.css" integrity="sha384-nn4HPE8lTHyVtfCBi5yW9d20FjT8BJwUXyWZT9InLYax14RDjBj46LmSztkmNP9w" crossorigin="anonymous">
    <link href="{{.CacheKey}}/static/css/styles.css" rel="stylesheet">
    {{range $i, $src := .data.DirectIncludes }}
      <script src="{{$src}}" type="text/javascript"></script>
    {{end}}
    <style>
	    {{ safeCss .data.InlineCss}}
    </style>
		{{.HtmlHead}}
  </head>
  <body>
    <div class="container">
      <div class="wrapper">

      {{if .data.DisplayHeader }}
        <div class="header">
          {{safeHtml .data.Header.ComponentHtml}}
        </div>
      {{end}}

      {{ if not .degraded }}
        {{ if .data.BodyTemplateColumns }}
          <div class="content" style="
          display:grid;
          grid-template-columns:{{safeCss .data.BodyTemplateColumns}};
          grid-template-rows:{{safeCss .data.BodyTemplateRows}};
          grid-row-gap:{{safeCss .data.BodyRowGap}};
          grid-column-gap:{{safeCss .data.BodyColumnGap}};
          justify-content:{{safeCss .data.JustifyContent}};
          align-content:{{safeCss .data.AlignContent}}">
            {{.LayoutContent}}
          </div>
        {{ else }}
          <div class="content">
            {{.LayoutContent}}
          </div>
        {{end}}
      {{ else }}
        <div class="content">
          {{.LayoutContent}}
        </div>
      {{ end }}
      {{if .data.DisplayFooter }}
        <div class="footer">
          {{safeHtml .data.Footer.ComponentHtml}}
        </div>
      {{end}}
      <div class="modals">
        {{range $index, $element := .data.ModalItems}}
          {{ modalHtml $element }}
        {{end}}
      </div>
      </div>
    </div>
  </body>
</html>
