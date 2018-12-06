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
      <form action="/" name=l method="GET">
        <h2>{{call .i18n "Choose_language"}}</h2>
          <p><label><input type="radio" name="lang" value="en-US">{{call .i18n "English"}}</label></p>
          <p><label><input type="radio" name="lang" value="ru-RU">{{call .i18n "Russian"}}</label></p>
          <input class="button" type="submit" onclick="l.action='/';" value="{{call .i18n "Submit_lang"}}">
      </form>
        <form action = "/best" name=f1 method = "GET">
          <h2>{{call .i18n "Choose_your_currency"}}</h2>
          <div class="warning">
            <p>{{.warningCurrency}}</p>
          </div>
          <p><label><input onClick="setAllCheckboxes('currency', this);" type="checkbox" />{{call .i18n "Select_All"}}</label></p>
          <div id="currency">
            <p><label><input type="checkbox" name="currency" value="usd">USD</label></p>
            <p><label><input type="checkbox" name="currency" value="eur">EUR</label></p>
          </div>
          <h2>{{call .i18n "Buy_or_Sale"}}</h2>
          <p><label><input type="radio" name="option" value="buy">{{call .i18n "Buy"}}</label></p>
          <p><label><input type="radio" name="option" value="sale">{{call .i18n "Sale"}}</label></p>
          <p><label><input type="radio" name="option" value="both" checked>{{call .i18n "Both"}}</label></p>
          <h2>{{call .i18n "Choose_your_banks"}}</h2>
          <div class="warning">
            <p>{{.warningBank}}</p>
          </div>
          <p><label><input onClick="setAllCheckboxes('banks', this);" type="checkbox" />{{call .i18n "Select_All"}}</label></p>
          <div id="banks">
            <p><label><input type="checkbox" name="bank" value="privat">{{call .i18n "Privat_Bank"}}</label></p>
            <p><label><input type="checkbox" name="bank" value="otp">{{call .i18n "OTP_Bank"}}</label></p>
            <p><label><input type="checkbox" name="bank" value="pireus">{{call .i18n "Pireus_Bank"}}</label></p>
          </div>
          <input class="button" type="submit" onclick="f1.action='/best';" value="{{call .i18n "Best_Choice"}}">      
          <input class="button" type="submit" onclick="f1.action='/comparision';" value="{{call .i18n "Compare"}}">
        </form>
      </div>
    </div>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
