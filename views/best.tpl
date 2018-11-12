<h1>Покупка</h1>
<ul>
{{range .Sale}}
<li>
Банк: {{.BankName}}<br>Валюта: {{.CodeAlpha}}<br>Покупка: {{.RateBuy}}<br>Продажа: {{.RateSale}}<br>
</li>
<br>
{{end}}
</ul>

<h1>Продажа</h1>
<ul>
{{range .Buy}}
<li>
Банк: {{.BankName}}<br>Валюта: {{.CodeAlpha}}<br>Покупка: {{.RateBuy}}<br>Продажа: {{.RateSale}}<br>
</li>
<br>
{{end}}
</ul>