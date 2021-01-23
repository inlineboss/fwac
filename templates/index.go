package templates

const Html string = `
<!DOCTYPE html>
<html lang="en">

<head>
    <title>DropIn</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css"
        integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"></script>
    <style>
        .icon-img {
            width: 40px;
            border-style: none;
        }
    </style>
</head>

<body>
    <header>
        <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
            <a class="navbar-brand" href="#">
                <h1>DropIn</h1>
            </a>
        </nav>

    </header>

    <ul class="nav nav-tabs">
        {{$last := len .Proxy.URL.Road}}
        {{ range $i, $e := .Proxy.URL.Road }}
            <li class="nav-item">
                {{ if eq (minus $last 1) $i}}
                <a class="nav-link active" href="{{ $e.Link }}"> <h3>{{ $e.Name }}</h3></a>
                {{ else }}
                <a class="nav-link" href="{{ $e.Link }}"> <h3>{{ $e.Name }}</h3></a>
                {{ end }}
            </li>
        {{ end }}
    </ul>

    {{ $folderImageUrl := "https://github.com/inlineboss/fwac/blob/master/templates/folder-icon.png?raw=true" }}
    {{ $fileImageUrl := "https://github.com/inlineboss/fwac/blob/master/templates/file-icon.png?raw=true" }}

    <ul class="list-group list-group-flush">
        {{ range .FileInfo }}
        <li class="list-group-item">
            <a href="{{ .Link }}">
                {{ if eq .File.Type "Folder"}}
                <img src="https://github.com/inlineboss/fwac/blob/master/templates/folder-icon.png?raw=true"
                    class="icon-img img-thumbnail">
                {{ else }}
                <img src="https://github.com/inlineboss/fwac/blob/master/templates/file-icon.png?raw=true"
                    class="icon-img img-thumbnail">
                {{ end }}
                {{ .File.Name }}
            </a>
        </li>
        {{ end }}
    </ul>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.4.1.slim.min.js"
        integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n"
        crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js"
        integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo"
        crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js"
        integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6"
        crossorigin="anonymous"></script>
</body>
</html>
`
