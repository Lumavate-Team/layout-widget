<!DOCTYPE html>
<html lang="en">
  <head>
  <script type="text/javascript" src="/ga.js?pageTitle={{.data.InstanceName}}"></script>
  {{ if eq .mode "KNOCKOUT" }}
    <script type='text/javascript' src='https://cdnjs.cloudflare.com/ajax/libs/knockout/3.5.0/knockout-min.js'></script>
  {{ end }}

  {{if .gtm }}
  <!-- Google Tag Manager -->
  <script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
  new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
  j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
  'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
  })(window,document,'script','dataLayer','{{.gtm}}');</script>
  <!-- End Google Tag Manager -->
  {{ end }}
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
    <link rel="manifest" href="/manifest.json" crossOrigin="use-credentials">
    <link rel="shortcut icon" href="/iot/favicon.ico">
    <meta name="apple-mobile-web-app-capable" content="no">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta property="og:image" content="{{.dnsInfo}}/iot/android-chrome-512x512.png" />

    <link href="{{.CacheKey}}/static/css/styles.css" rel="stylesheet">
    <link rel="stylesheet" type="text/css" href="{{.CacheKey}}/static/css/addtohomescreen.css">
    {{range $prop := .data.StyleData }}
      {{if hasSuffix $prop.Name "FontFamily"}}
        <link href="https://fonts.googleapis.com/css?family={{ $prop.Value }}" rel="stylesheet">
      {{end}}
    {{end}}

    {{range $i, $href := .data.DirectCssIncludes }}
      <link href="{{$href}}" rel="stylesheet">
    {{end}}

    {{range $i, $src := .data.DirectIncludes }}
      <script src="{{$src}}" type="text/javascript"></script>
    {{end}}

    <script id="aths" async type="text/javascript" src="{{.CacheKey}}/static/js/addtohomescreen.js"></script>
    <script id="luma-core" type="text/javascript" src="{{.CacheKey}}/core/luma-core.js"></script>


    <script>
      var lightenColor = function(color, percent) {
          color = color.replace('#', '');
          var num = parseInt(color,16),
          amt = Math.round(2.55 * percent),
          R = (num >> 16) + amt,
          B = (num >> 8 & 0x00FF) + amt,
          G = (num & 0x0000FF) + amt;

          return '#' + (0x1000000 + (R<255?R<1?0:R:255)*0x10000 + (B<255?B<1?0:B:255)*0x100 + (G<255?G<1?0:G:255)).toString(16).slice(1);
      };
      document.addEventListener("DOMContentLoaded", function(event) {
        body = document.querySelector('body');
        {{range $prop := .data.StyleData }}
          var prop_name = '--{{ $prop.Name }}'.replace(/\.?([A-Z])/g, function (x,y){return "-" + y.toLowerCase()}).replace(/^_/, "");
          {{if hasSuffix $prop.Name "ColorFamily"}}
            a = 100;
            while (a > 0) {
              body.style.setProperty(prop_name + '-' + a, lightenColor('{{ $prop.Value }}', 100-a));
              a = a - 10;
            }
          {{end}}
          body.style.setProperty(prop_name, '{{ $prop.Value }}');
        {{end}}
      });
    </script>

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

{{ if .gtm }}
  <!-- Google Tag Manager (noscript) -->
<noscript><iframe src="https://www.googletagmanager.com/ns.html?id={{.gtm}}"
height="0" width="0" style="display:none;visibility:hidden"></iframe></noscript>
<!-- End Google Tag Manager (noscript) -->
{{ end }}

    {{range $index, $element := .data.LogicItems}}
      {{if eq $element.ComponentData.Placement "top"}}
        {{ logicHtml $element }}
      {{end}}
    {{end}}

    <div class="container">
      <div class="wrapper">
      <script>
        function getCookieValue(a) {
          var b = document.cookie.match('(^|;)\\s*' + a + '\\s*=\\s*([^;]+)');
          return b ? b.pop() : '';
        }

        function getAuthUrl() {
          let token = getCookieValue("pwa_jwt")
          token = token.split(".")[1];
          let decodedToken = JSON.parse(atob(token));

          return decodedToken["authUrl"]
        }

        function getActivationId(){
          let token = getCookieValue("pwa_jwt")
          token = token.split(".")[1];
          let decodedToken = JSON.parse(atob(token));

          return decodedToken["activationId"]
        }

        function getSingleUseToken(onSuccess, onNoAuth, onError) {
          let xsrf, xsrflist;
          xsrf = getCookieValue("_xsrf");
          xsrflist = xsrf.split("|");
          let xhr = new XMLHttpRequest();
          xhr.responseType = 'json';

          xhr.open('POST', window.location.href.split('?')[0] + '/sut');
          xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
          xhr.onload = function() {
            if (xhr.status === 200) {
              if (onSuccess) {
                let token = xhr.response.alg + " " + xhr.response.token + " UrlRef=" + xhr.response.urlRef
                onSuccess(token);
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

        if(/\?mode=degraded/.test(window.location.search) === false){
          if(window.CSS === undefined || window.CSS.supports === undefined){
            window.location.href = window.location.href + '?mode=degraded'
          }

          if(window.CSS.supports('display', 'grid') === false){
            window.location.href = window.location.href + '?mode=degraded'
          }
        }
      </script>
      {{if .data.DisplayHeader }}
        <div class="header">
          {{safeHtml .data.Header.ComponentHtml}}
        </div>
      {{end}}

      {{ if eq .mode "CSSGRID"}}
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
      {{ end }}
      {{ if eq .mode "KNOCKOUT" }}
        {{ .data.ViewTemplate }}
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
  {{range $index, $element := .data.LogicItems}}
    {{if eq $element.ComponentData.Placement "bottom"}}
      {{ logicHtml $element }}
    {{end}}
  {{end}}

  {{range $index, $element := .resources.Pages}}
      <luma-core-page id='{{$element.Id}}' url='{{$element.Url}}'></luma-core-page>
  {{end}}

  {{range $index, $element := .resources.Microservices}}
      <luma-core-microservice id='{{$element.Id}}' uri='{{$element.Url}}'></luma-core-microservice>
  {{end}}

  <luma-core-context></luma-core-context>
  <script>
    {{ if eq .mode "KNOCKOUT" }}
      window.strings = {}
      {{range $i, $string := .data.Translations }}
        window.strings['{{ $string.ComponentData.StringId }}'] = '{{ $string.ComponentData.String }}';{{end}}

      window.variables = {}
      {{range $i, $variable := .data.Variables }}
        window.variables['{{ $variable.ComponentData.VariableId }}'] = '{{ $variable.ComponentData.Variable }}';{{end}}

      import_strings = function(o) {
        for (k in window.strings) {
          o[k] = window.strings[k];
        }
      }

      import_variables = function(o) {
        for (k in window.variables) {
          o[k] = window.variables[k];
        }
      }
    {{ end }}

    var lc = document.querySelector('luma-core-context');
    lc.componentOnReady().then(function() {
      lc.authData = {{ .auth_json }};
      lc.activationData = {{ .activation_json }};
      lc.domainData = {{ .domain_json }};
    });
    {{ .data.Script }}
    {{ .data.ViewModel }}

    orig = ko.applyBindings
    ko.applyBindings = function(arg1) {
      import_strings(arg1);
      import_variables(arg1);
      orig(arg1)
    }
  </script>
  <script type="text/javascript">
    var athsScript = document.querySelector('#aths');
    athsScript.addEventListener('load', function() {
      var userAgent = navigator.userAgent.toLowerCase();
      var isAndroid = userAgent.indexOf('android') > -1;

      if((!isAndroid) && {{.data.HomeScreen.ShowAddToHome}} == true) {
        var HomeScreenConfig = {
          appID: 'lumavate.addtohomescreen.' + getActivationId(),
          skipFirstVisit: {{.data.HomeScreen.SkipFirst}},
          startDelay: {{.data.HomeScreen.StartDelay}},
          lifespan: {{.data.HomeScreen.Lifespan}},
          maxDisplayCount: {{.data.HomeScreen.DisplayCount}},
          message: {{.data.HomeScreen.Message}}
        };
        addToHomescreen(HomeScreenConfig);
      };
    });
  </script>
  <script type="text/javascript" src="/iot/sw-register.min.js"></script>
  </body>
</html>
