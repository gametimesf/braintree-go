package braintree

type PaymentMethod interface {
	GetCustomerId() string
	GetToken() string
	IsDefault() bool
	GetImageURL() string
}

type PaymentMethodMerchantResult struct {
	Nonce string `xml:"nonce"`
}

type PaymentMethodMerchantGrant struct {
	PaymentMethodResult PaymentMethodMerchantResult `xml:"payment_method_nonce"`
}
