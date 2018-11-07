<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  

  <div style= "display: flex">
  <div style = "border: 3px solid #1d1d1d; padding: 10px">
  <form action = "/best" name=f1 method = "GET">
    <h2>Choose your currency</h2>
      <p>
        <input type="checkbox" name="currency" value="usd">USD <br/>
        <input type="checkbox" name="currency" value="eur">EUR
      </p>
      <p> 
        <input type="radio" name="option" value="buy"> Buy<br/>
        <input type="radio" name="option" value="sale"> Sale<br/>
        <input type="radio" name="option" value="both" checked> Both<br/>
      </p>
      <p>
        <input type="checkbox" name="bank" value="privat">Privat Bank<br/>
        <input type="checkbox" name="bank" value="otp">OTP Bank<br/>
        <input type="checkbox" name="bank" value="pireus">Pireus Bank<br/>
      </p>
      <input type="submit" value="Best choice">      
      <input type="submit" onclick="f1.action='/comparision';" value="Compare">
  
    </form>
  </div>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
