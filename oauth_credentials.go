package braintree

import "time"

type OAuthCredentialsResult struct {
	AccessToken  string    `xml:"access_token"`
	RefreshToken string    `xml:"refresh_token"`
	ExpiresAt    time.Time `xml:"expires_at"`
}

type OAuthCredentials struct {
	Credentials OAuthCredentialsResult `xml:"credentials"`
}

type OAuthRevokeResult struct {
	// I cannot find the type for the success attribute.
	// It is not in the documentation, it is not in the code.
	// I seems as long as it is not exist, it is a success.
	Success interface{} `xml:"success"`
}
type OAuthRevokeToken struct {
	Result OAuthRevokeResult `xml:"result"`
}
