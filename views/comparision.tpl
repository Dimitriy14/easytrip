<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
    <link rel="stylesheet" href="../static/css/comparision.css">
</head>
<body>
    <div class = "mycontainer">
        <h1 class="page-name">Rates rating</h1>
        {{range $key, $val := .Banks}}
            <div class="bank">
                <h2 class="bank-name">{{$val.BankName}}</h2>
                <ul class="bank-data">
                    <li>Code Alpha: {{$val.CodeAlpha}}</li>
                    <li>Rate Buy: {{$val.RateBuy}}</li>
                    <li>Rate Sale: {{$val.RateSale}}</li>
                </ul>
            </div>
        {{end}}
    </div>
    
</body>
</html>