<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="NikoNiko Calendar">
    <meta name="author" content="Thibault NORMAND">
    <meta name="_xsrf" content="{{.xsrf_token}}">
    <title>{{if .Title}}{{.Title}} | {{end}}NikoNiko</title>

    <!-- bulma -->
    <link href="//cdnjs.cloudflare.com/ajax/libs/bulma/0.3.1/css/bulma.min.css" rel="stylesheet">
    <!-- Font Awesome -->
    <link href="//maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">

    <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
        <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
</head>
<body>

    <nav class="nav has-shadow">
      <div class="container">
        <div class="nav-left">
          <a href="/" class="nav-item">
            <span class="title">NikoNiko</span>
          </a>
        </div>
        <div class="nav-toggle">
          <span></span>
          <span></span>
          <span></span>
        </div>
        <div class="nav-right nav-menu">
        </div>
      </div>
    </nav>

    {{if .flash_err}}<div class="notification is-danger is-marginless">{{.flash_err}}</div>{{end}}
    {{if .flash_warn}}<div class="notification is-warning is-marginless">{{.flash_warn}}</div>{{end}}
    {{if .flash_info}}<div class="notification is-info is-marginless">{{.flash_info}}</div>{{end}}

    {{yield}}

</body>
</html>
