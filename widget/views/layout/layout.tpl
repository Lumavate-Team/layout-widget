<!DOCTYPE html>
<html>
  <head>
    <title>
    </title>
    <meta charset="utf-8">
    <base href="{{.WidgetUrlPrefix}}">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <noscript>Javascript required for this site to work.</noscript>

    <link rel="apple-touch-icon" sizes="180x180" href="/iot/favicon-180x180.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/iot/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/iot/favicon-16x16.png">
    <link rel="manifest" href="/manifest.json">
    <link rel="shortcut icon" href="/iot/favicon.ico">
    <meta name="apple-mobile-web-app-capable" content="no">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta property="og:image" content="{{.dnsInfo}}/iot/favicon-180x180.png" />

    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Teko:400,500" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro" rel="stylesheet">
    <link rel="stylesheet" href="https://unpkg.com/purecss@1.0.0/build/pure-min.css" integrity="sha384-nn4HPE8lTHyVtfCBi5yW9d20FjT8BJwUXyWZT9InLYax14RDjBj46LmSztkmNP9w" crossorigin="anonymous">  
		<style>
			#tiles {
				position:absolute;
				height:100%;
				width:100%;
				display:grid;
				grid-template-columns: {{.data.GridTemplateColumns}};
				grid-template-rows: {{.data.GridTemplateRows}};
				grid-gap: {{ .data.Padding }}px;
			}
		</style>
		{{.HtmlHead}}
  </head>
  <body>
  <div class="body-wrapper">
    {{.HeaderContent}}
        {{.LayoutContent}}

    </div>
    {{.FooterContent}}
    <script type="text/javascript" src="https://code.jquery.com/jquery-2.0.3.min.js"></script>

    <script type="text/javascript" src="lc/lumavate-components.js"></script>
    {{.Scripts}}

    <script type="text/javascript">
      $(function(){

        var setWrapperHeight = function(){
          $(".body-wrapper").css("height", window.innerHeight-58);
        }

        var resizeTimer;
        $(window).on("resize orientationchange", function(e){
          clearTimeout(resizeTimer);
          resizeTimer = setTimeout(function() {

            setWrapperHeight();
          },250);
        });
        setWrapperHeight();
      });
    </script>
  </body>
</html>
