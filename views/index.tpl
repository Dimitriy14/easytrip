<!DOCTYPE html>

<html>
<head>
  <title>Exrates</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="../static/css/main.css">
</head>

<body>
  <div class="container">
    <div class="form-window">
      <form action = "/best" name=f1 method = "GET">
        <div class="currency-block">
          <h2>Choose your currency</h2>
          <p><input type="checkbox" name="currency" value="usd">USD</p>
          <p><input type="checkbox" name="currency" value="eur">EUR</p>
        </div>
        <h2>Buy or Sale</h2>
        <p><input type="radio" name="option" value="buy"> Buy</p>
        <p><input type="radio" name="option" value="sale"> Sale</p>
        <p><input type="radio" name="option" value="both" checked> Both</p>
        <h2>Choose your banks</h2>
        <p><input type="checkbox" name="bank" value="privat">Privat Bank</p>
        <p><input type="checkbox" name="bank" value="otp">OTP Bank</p>
        <p><input type="checkbox" name="bank" value="pireus">Pireus Bank</p>
        <input type="submit" value="Best choice">      
        <input type="submit" onclick="f1.action='/comparision';" value="Compare">
      </form>
    </div>
  </div>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
