package braintree

import (
	"context"
	"encoding/xml"
)

type OAuthGateway struct {
	*Braintree
}

type (
	OAuthCreateTokenFromCodeRequest struct {
		XMLName   xml.Name `xml:"credentials"`
		Code      string   `xml:"code"`
		GrantType string   `xml:"grant_type"`
	}

	OAuthRefreshCredentialsRequest struct {
		XMLName      xml.Name `xml:"credentials"`
		RefreshToken string   `xml:"refresh_token"`
		GrantType    string   `xml:"grant_type"`
	}
)

// CreateTokenFromCode creates an OAuth credentials from an authorization code.
func (g *OAuthGateway) CreateTokenFromCode(ctx context.Context, code string) (*OAuthCredentials, error) {
	req := OAuthCreateTokenFromCodeRequest{Code: code, GrantType: "authorization_code"}

	resp, err := g.executeBaseVersion(ctx, "POST", "oauth/access_tokens", &req, apiVersion6)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
		return resp.oauthCredentials()
	}

	return nil, &invalidResponseError{resp}
}

// CreateTokenFromRefreshToken creates an OAuth credentials from a refresh token.
func (g *OAuthGateway) CreateTokenFromRefreshToken(ctx context.Context, refreshToken string) (*OAuthCredentials, error) {
	req := OAuthRefreshCredentialsRequest{RefreshToken: refreshToken, GrantType: "refresh_token"}

	resp, err := g.executeBaseVersion(ctx, "POST", "oauth/access_tokens", &req, apiVersion6)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
		return resp.oauthCredentials()
	}
	return nil, &invalidResponseError{resp}
}
