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

    <link href="{{.CacheKey}}/static/css/styles.css" rel="stylesheet">
    {{range $i, $href := .data.DirectCssIncludes }}
      <link href="{{$href}}" rel="stylesheet">
    {{end}}

    {{range $i, $src := .data.DirectIncludes }}
      <script src="{{$src}}" type="text/javascript"></script>
    {{end}}
    <style>
      body {
        {{ if .data.DisplayBackgroundImage }}
          background-image: url({{.data.BackgroundImage.Preview}});
          background-repeat:repeat;
        {{ else if .data.BackgroundColor }}
          background-color: {{.data.BackgroundColor}};
        {{ end }}
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="wrapper">
      <script>
        function getCookieValue(a) {
          var b = document.cookie.match('(^|;)\\s*' + a + '\\s*=\\s*([^;]+)');
          return b ? b.pop() : '';
        }

        function getSingleUseToken(onSuccess, onNoAuth, onError) {
          var xsrf, xsrflist;
          xsrf = getCookieValue("_xsrf");
          xsrflist = xsrf.split("|");
          xhr = new XMLHttpRequest();
          xhr.responseType = 'json';

          xhr.open('POST', window.location.href.split('?')[0] + '/single-use-token');
          xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
          xhr.onload = function() {
            if (xhr.status === 200) {
              if (onSuccess) {
                onSuccess(xhr.response.token);
              }
            }
            else if (xhr.status === 500) {
              if (onError) {
                onError();
              } else {
                console.log('Error requesting single-use-token');
              }
            } else {
              if (onNoAuth) {
                onNoAuth();
              } else {
                console.log('Not authorized to request single-use-token');
              }
            }
          };
          xhr.send('_xsrf=' + atob(xsrflist[0]));
        }

      </script>
      {{if .data.DisplayHeader }}
        <div class="header">
          {{safeHtml .data.Header.ComponentHtml}}
        </div>
      {{end}}

      {{ if not .degraded }}
        {{ if eq .data.BodyProperties.ComponentType "body-items-advanced" }}
          <div class="content" style="display:grid;
            grid-template-columns:{{safeCss .data.BodyProperties.ComponentData.BodyTemplateColumns}};
            grid-template-rows:{{safeCss .data.BodyProperties.ComponentData.BodyTemplateRows}};
            grid-row-gap:{{safeCss .data.BodyProperties.ComponentData.BodyRowGap}};
            grid-column-gap:{{safeCss .data.BodyProperties.ComponentData.BodyColumnGap}};
            justify-content:{{safeCss .data.BodyProperties.ComponentData.JustifyContent}};
            align-content:{{safeCss .data.BodyProperties.ComponentData.AlignContent}}">
        {{ else }}
          <div class="content" style="display:grid;
            grid-template-columns:{{safeCss .data.BodyProperties.ComponentData.BodyTemplateColumns}};
            grid-template-rows:{{safeCss .data.BodyProperties.ComponentData.BodyTemplateRows}};
            max-width: {{safeCss .data.BodyProperties.ComponentData.BodyMaxWidthStr}}">
        {{ end }}
          {{.LayoutContent}}
        </div>
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
