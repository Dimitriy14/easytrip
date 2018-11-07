<h1>Hello</h1>

{{.re}}

{{range $key, $val := .Banks}}
    <div class="bank" style="border: 3px solid #1d1d1d; margin: 5px">
        <h2 class="bank-name">{{$val.BankName}}</h2>
        <ul class="bank-data">
            <li>Code Alpha: {{$val.CodeAlpha}}</li>
            <li>Rate Buy: {{$val.RateBuy}}</li>
            <li>Rate Sale: {{$val.RateSale}}</li>
        </ul>
    </div>
{{end}}