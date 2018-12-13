<<<<<<< HEAD
    <div class = "mycontainer">
        <h1 class="page-name">Rates rating</h1>
=======
<div class = "mycontainer">
        <h1 class="page-name">{{call .i18n "Rates_rating"}}</h1>
        {{if $self := .}}
>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
        {{range $key, $val := .Banks}}
            <div class="bank">
                <h2 class="bank-name">{{call $self.i18n $val.BankName}}</h2>
                <ul class="bank-data">
                    <li>{{call $self.i18n "Code_Alpha"}}: {{$val.CodeAlpha}}</li>
                    <li>{{call $self.i18n "Rate_Buy"}}: {{$val.RateBuy}}</li>
                    <li>{{call $self.i18n "Rate_Sale"}}: {{$val.RateSale}}</li>
                </ul>
            </div>
        {{end}}
        {{end}}
        <a href="/" class="back">{{call .i18n "Back"}}</a>
    </div>
