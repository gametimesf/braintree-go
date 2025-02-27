package braintree

type credentials interface {
	Environment() Environment
	GraphQLEnvironment() Environment
	MerchantID() string
	AuthorizationHeader() string
	AuthorizationHeaderWithClientCreds() string
}
