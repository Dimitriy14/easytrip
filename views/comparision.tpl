<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
    <link href="/static/css/comparision.css" rel="stylesheet">
</head>
<body>
    <div class = "mycontainer">
     {{.flash.error}}
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
<<<<<<< HEAD
        <a href = "/" class="button">Back to main</a>
=======
        <a href="/" class="back">Back</a>
>>>>>>> dbd7f9e20e2df2bb8786b76473fcf66688bd2e77
    </div>
</body>
</html>