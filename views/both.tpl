<!DOCTYPE html>
<html>
<head>
    <title>Best</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link href="/static/css/main.css" rel="stylesheet">
</head>
<body>
<div class="bestBank">
<div class="mycontainer">
<div class="form-window">
<h1>Best sale</h1>
<ul>
{{range .Sale}}
<li>
Bank: {{.BankName}}<br>Currency: {{.CodeAlpha}}<br>Sale: {{.RateSale}}<br>
</li>
<br>
{{end}}
</ul>
</div>
</div>
</div>
<br>
<div class="bestBank">
<div class = "mycontainer">
<div class="form-window">
<h1>Best buy</h1>
<ul>
{{range .Buy}}
<li>
Bank: {{.BankName}}<br>Currency: {{.CodeAlpha}}<br>Buy: {{.RateBuy}}<br>
</li>
<br>
{{end}}
</ul>
<a href="/" class="back">Back</a>
</div>
</div>
</div>
</body>
</html>
