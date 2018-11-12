<!DOCTYPE html>
<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link href="/static/css/main.css" rel="stylesheet">
</head>
<body>

{{if or (eq .req "buy") (eq .req "both")}}
<div class="bestBank">
<h1>Покупка</h1>
<ul>
{{range .Buy}}
<li>
Банк: {{.BankName}}<br>Валюта: {{.CodeAlpha}}<br>Покупка: {{.RateBuy}}<br>Продажа: {{.RateSale}}<br>
</li>
<br>
{{end}}
</ul>
</div>
{{end}}

<br>

{{if or (eq .req "sale") (eq .req "both")}}
<div class="bestBank">
<h1>Продажа</h1>
<ul>
{{range .Sale}}
<li>
Банк: {{.BankName}}<br>Валюта: {{.CodeAlpha}}<br>Покупка: {{.RateBuy}}<br>Продажа: {{.RateSale}}<br>
</li>
<br>
{{end}}
</ul>
</div>
{{end}}
</body>
</html>
