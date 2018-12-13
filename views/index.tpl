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
          <p><label><input type="radio" name="lang" value="ua-UA">{{call .i18n "Ukrainian"}}</label></p>
          <div class="button-select">
          <input class="button" type="submit" onclick="l.action='/';" value="{{call .i18n "Submit_lang"}}">
          </div>
      </form>
      </div>

      <div class="form-window">
        <form action = "/best" name=f1 method = "GET">
          <div id="currency">
<<<<<<< HEAD
            <h2>Choose your currency</h2>
            <p><label><input onClick="setAllCheckboxes('currency', this);" type="checkbox" />Select All</label></p>
            {{if .IncorrectCurrency}}<div class="error">Choose currency</div>{{end}}
=======
            <h2>{{call .i18n "Choose_your_currency"}}</h2>
            <div class="warning">
              <p>{{if .isWarnCurr}} {{call .i18n .warningCurrency}}{{end}}</p>
            </div>
            <p><label><input onClick="setAllCheckboxes('currency', this);" type="checkbox" />{{call .i18n "Select_All"}}</label></p>
>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
            <p><label><input type="checkbox" name="currency" value="usd">USD</label></p>
            <p><label><input type="checkbox" name="currency" value="eur">EUR</label></p>
          </div>
          <div class="option">
<<<<<<< HEAD
            <h2>Buy or Sale</h2>
            <p><label><input type="radio" name="option" value="buy">Buy</label></p>
            <p><label><input type="radio" name="option" value="sale">Sale</label></p>
            <p><label><input type="radio" name="option" value="both" checked>Both</label></p>
          </div>
          <div id="banks">
            <h2>Choose your banks</h2>
            {{if .IncorrectBank}}<div class="error">Choose bank</div>{{end}}
            <p><label><input onClick="setAllCheckboxes('banks', this);" type="checkbox" />Select All</label></p>
            <p><label><input type="checkbox" name="bank" value="privat">Privat Bank</label></p>
            <p><label><input type="checkbox" name="bank" value="otp">OTP Bank</label></p>
            <p><label><input type="checkbox" name="bank" value="pireus">Pireus Bank</label></p>
=======
            <h2>{{call .i18n "Buy_or_Sale"}}</h2>
            <p><label><input type="radio" name="option" value="buy">{{call .i18n "Buy"}}</label></p>
            <p><label><input type="radio" name="option" value="sale">{{call .i18n "Sale"}}</label></p>
            <p><label><input type="radio" name="option" value="both" checked>{{call .i18n "Both"}}</label></p>
          </div>
          <div id="banks">
            <h2>{{call .i18n "Choose_your_banks"}}</h2>
            <div class="warning">
              <p>{{if .isWarnBank}} {{call .i18n .warningBank}}{{end}}</p>
            </div>
            <p><label><input onClick="setAllCheckboxes('banks', this);" type="checkbox" />{{call .i18n "Select_All"}}</label></p>
            <p><label><input type="checkbox" name="bank" value="privat">{{call .i18n "Privat_Bank"}}</label></p>
            <p><label><input type="checkbox" name="bank" value="otp">{{call .i18n "OTP_Bank"}}</label></p>
            <p><label><input type="checkbox" name="bank" value="pireus">{{call .i18n "Pireus_Bank"}}</label></p>
>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
          </div>
          <input class="button" type="submit" onclick="f1.action='/best';" value="{{call .i18n "Best_Choice"}}">      
          <input class="button" type="submit" onclick="f1.action='/comparision';" value="{{call .i18n "Compare"}}">
        </form>
<<<<<<< HEAD
=======

      </div>
      <div class="form-window">
          <div class="button-center"><a class="button" href = "/statistics">{{call .i18n "Get_statistics"}}</a></div>
>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
      </div>     
    </div>
    <div class="button-center"><a class="button" href = "/statistics">Get statistics</a></div>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
