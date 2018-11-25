<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link href="/static/css/main.css" rel="stylesheet">
</head>

<body>
    <div class="mycontainer">
      <div class="form-window">
        <form action = "/best" name=f1 method = "GET">
<<<<<<< HEAD
=======
          <h2>Choose your currency</h2>
          <div class="warning">
            <p>{{.warningCurrency}}</p>
          </div>
          <p><label><input onClick="setAllCheckboxes('currency', this);" type="checkbox" />Select All</label></p>
>>>>>>> cdafc8e1a6224d8b56d6ab90e36345f3820e947a
          <div id="currency">
            <h2>Choose your currency</h2>
            <p><label><input onClick="setAllCheckboxes('currency', this);" type="checkbox" />Select All</label></p>
            {{if .IncorrectCurrency}}<div class="error">Choose currency</div>{{end}}
            <p><label><input type="checkbox" name="currency" value="usd">USD</label></p>
            <p><label><input type="checkbox" name="currency" value="eur">EUR</label></p>
          </div>
<<<<<<< HEAD
          <div class="option">
            <h2>Buy or Sale</h2>
            <p><label><input type="radio" name="option" value="buy">Buy</label></p>
            <p><label><input type="radio" name="option" value="sale">Sale</label></p>
            <p><label><input type="radio" name="option" value="both" checked>Both</label></p>
          </div>
=======
          <h2>Buy or Sale</h2>
          <p><label><input type="radio" name="option" value="buy">Buy</label></p>
          <p><label><input type="radio" name="option" value="sale">Sale</label></p>
          <p><label><input type="radio" name="option" value="both" checked>Both</label></p>
          <h2>Choose your banks</h2>
          <div class="warning">
            <p>{{.warningBank}}</p>
          </div>
          <p><label><input onClick="setAllCheckboxes('banks', this);" type="checkbox" />Select All</label></p>
>>>>>>> cdafc8e1a6224d8b56d6ab90e36345f3820e947a
          <div id="banks">
            <h2>Choose your banks</h2>
            {{if .IncorrectBank}}<div class="error">Choose bank</div>{{end}}
            <p><label><input onClick="setAllCheckboxes('banks', this);" type="checkbox" />Select All</label></p>
            <p><label><input type="checkbox" name="bank" value="privat">Privat Bank</label></p>
            <p><label><input type="checkbox" name="bank" value="otp">OTP Bank</label></p>
            <p><label><input type="checkbox" name="bank" value="pireus">Pireus Bank</label></p>
          </div>
          <input class="button" type="submit" onclick="f1.action='/best';" value="Best choice">      
          <input class="button" type="submit" onclick="f1.action='/comparision';" value="Compare">
        </form>
      </div>
    </div>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
