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
      <h1>Best buy</h1>
      <ul>
{{range .Buy}}
        <li>Bank: {{.BankName}}<br>Currency: {{.CodeAlpha}}<br>Buy: {{.RateBuy}}
        </li>
<br>
{{end}}
      </ul>
    </div>
{{end}}

{{if or (eq .req "sale") (eq .req "both")}}
    <div class="bestBank">
<h1>Best sale</h1>
      <ul>
{{range .Sale}}
        <li>Bank: {{.BankName}}</li>
         <li>Currency: {{.CodeAlpha}}</li>
         <li>Sale: {{.RateSale}}</li>
<br>
{{end}}
      </ul>
    </div>
{{end}}
</body>
</html>
