{{range .}}
<li>
Банк: {{.BankName}}<br>Валюта: {{.CodeAlpha}}<br>Покупка: {{.RateBuy}}<br>Продажа: {{.RateSale}}
</li>
{{end}}