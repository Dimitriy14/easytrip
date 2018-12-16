<div class = "mycontainer">

<<<<<<< HEAD
<h1>Best Banks</h1>
<h2>{{.TitleBuy}}</h2>
{{range .Buy}}
<div class="bank">
<ul>
<h2 class="bank-name">{{.BankName}}</h2>
=======
<h1>{{call .i18n "Best_Banks"}}</h1>
<h2>{{call .i18n .TitleBuy }}</h2>
{{if $self := .}}
{{range .Buy}}
<div class="bank">
<ul>
<h2 class="bank-name">{{call $self.i18n .BankName}}</h2>
>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
<li>
{{.CodeAlpha}}: {{.RateBuy}}
</li>
<br>
</ul>
</div>
{{end}}
<<<<<<< HEAD

<h2>{{.TitleSale}}</h1>
{{range .Sale}}
<div class="bank">
<ul>
<h2 class="bank-name">{{.BankName}}</h2>
=======
{{end}}

<h2>{{call .i18n .TitleSale }}</h1>
{{if $self := .}}
{{range .Sale}}
<div class="bank">
<ul>
<h2 class="bank-name">{{call $self.i18n .BankName}}</h2>
>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
<li>
{{.CodeAlpha}}: {{.RateSale}}
</li>
<br>
</ul>
</div>
{{end}}
<<<<<<< HEAD
<a href="/" class="back">Back</a>
=======
{{end}}
<a href="/" class="back">{{call .i18n "Back"}}</a>
>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
</div>