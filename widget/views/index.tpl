<!DOCTYPE html>

<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">

  <link rel="apple-touch-icon" sizes="180x180" href="/iot/favicon-180x180.png">
  <link rel="icon" type="image/png" sizes="32x32" href="/iot/favicon-32x32.png">
  <link rel="icon" type="image/png" sizes="16x16" href="/iot/favicon-16x16.png">
  <link rel="manifest" href="/manifest.json">
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-status-bar-style" content="black">


  <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
  <link href="https://fonts.googleapis.com/css?family=Teko:400,500" rel="stylesheet">

  <link href="static/css/styles.css" rel="stylesheet">
</head>

<body style="background-color: {{.data.BackgroundColor}}">
  <header>
    <div class='header-container'>
        {{ componentHtml .data.Title }}
    </div>
  </header>
    <div class="zoom-img">
      <img src= {{.image}}>
    </div>
  <footer>
    {{ componentHtml .data.NavBar }}
  </footer>

  <script src="lc/lumavate-components.js"></script>
  <script src="ims/ims-lumavate-components.js"></script>
</body>
</html>
