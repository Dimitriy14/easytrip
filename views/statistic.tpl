<!DOCTYPE html>
<html>
<head>
    <title>Statistic</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link href="/static/css/main.css" rel="stylesheet">
</head>
<body>
<h2 class = "stathead">{{call .i18n "Get_statistics"}}</h2>
    <div class="block-container">
        {{range $_, $val := .Stat}}
            <div class = "form-window">
                <ul>
                    {{range $key,$stat := $val}}
                        <li>{{$stat}}</li>
                    {{end}}
                </ul>   
            </div>  
        {{end}}
    </div> 
    <div class="button-center"><a href="/" class="back">{{call .i18n "Back"}}</a></div>
</body>
</html>