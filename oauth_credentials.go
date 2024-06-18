package braintree

import "time"

type OAuthCredentials struct {
	XMLName      string    `xml:"credentials"`
	AccessToken  string    `xml:"access-token"`
	RefreshToken string    `xml:"refresh-token"`
	ExpiresAt    time.Time `xml:"expires-at"`
	TokenType    string    `xml:"token-type"`
	Scope        string    `xml:"scope"`
}
