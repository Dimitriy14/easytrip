package clients

import (
	"io/ioutil"
	"net/http"
)

// BankUAClient represents a common client to interact with a remote Bank Service
type BankUAClient interface {
	Get() (body []byte, err error)
}

// BankUAClientImpl implements BankUAClient interface
type BankUAClientImpl struct {
	baseURL    string
	httpClient *http.Client
}

// Get returns a remote Bank Service response
func Get() (body []byte, err error) {
	res, err := httpClient.Get(baseURL)
	if err != nil {
		return
	}

	defer func() {
		if res.Body != nil {
			res.Body.Close()
		}
	}()

	return ioutil.ReadAll(res.Body)
}

// New creates a new BankUAClient instance
func New() BankUAClientImpl {
	return BankUAClientImpl{
		baseURL:    "http://bank-ua.com/export/exchange_rate_cash.json",
		httpClient: &http.Client{},
	}
}
