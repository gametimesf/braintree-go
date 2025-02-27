package braintree

import "fmt"

type Environment struct {
	baseURL string
}

func NewEnvironment(baseURL string) Environment {
	return Environment{baseURL: baseURL}
}

func (e Environment) BaseURL() string {
	return e.baseURL
}

var (
	Development = NewEnvironment("http://localhost:3000")
	Sandbox     = NewEnvironment("https://api.sandbox.braintreegateway.com:443")
	Production  = NewEnvironment("https://api.braintreegateway.com:443")

	GraphQLDevelopment = NewEnvironment("http://localhost:3000")
	GraphQLSandbox     = NewEnvironment("https://payments.sandbox.braintree-api.com")
	GraphQLProduction  = NewEnvironment("https://payments.braintree-api.com")
)

func EnvironmentFromName(name string) (Environment, error) {
	switch name {
	case "development":
		return Development, nil
	case "sandbox":
		return Sandbox, nil
	case "production":
		return Production, nil
	}
	return Environment{}, fmt.Errorf("unknown environment %q", name)
}

func GraphQLEnvironmentFromName(name string) (Environment, error) {
	switch name {
	case "development":
		return GraphQLDevelopment, nil
	case "sandbox":
		return GraphQLSandbox, nil
	case "production":
		return GraphQLProduction, nil
	}
	return Environment{}, fmt.Errorf("unknown environment %q", name)
}
