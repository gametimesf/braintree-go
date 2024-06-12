package braintree

import "context"

type OAuthGateway struct {
	*Braintree
}

type (
	OAuthCreateTokenFromCodeCredentials struct {
		Code string `xml:"code"`
		// grant-type will be set to "authorization_code"
		GrantType string `xml:"grant_type"`
	}

	OAuthCreateTokenFromCodeRequest struct {
		Credentials OAuthCreateTokenFromCodeCredentials `xml:"credentials"`
	}

	OAuthCreateTokenFromRefreshTokenCredentials struct {
		RefreshToken string `xml:"refresh_token"`

		// grant-type will be set to "refresh_token"
		GrantType string `xml:"grant_type"`
	}

	OAuthCreateTokenFromRefreshTokenRequest struct {
		Credentials OAuthCreateTokenFromRefreshTokenCredentials `xml:"credentials"`
	}

	OAuthRevokeAccessTokenRequest struct {
		AccessToken string `xml:"token"`
	}
)

func (g *OAuthGateway) CreateTokenFromCode(ctx context.Context, request *OAuthCreateTokenFromCodeRequest) (*OAuthCredentials, error) {
	request.Credentials.GrantType = "authorization_code"

	resp, err := g.executeVersion(ctx, "POST", "oauth/access_tokens", request, apiVersion6)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.oauthCredentials()
	}
	return nil, &invalidResponseError{resp}
}

func (g *OAuthGateway) CreateTokenFromRefreshToken(ctx context.Context, request *OAuthCreateTokenFromRefreshTokenRequest) (*OAuthCredentials, error) {
	request.Credentials.GrantType = "refresh_token"

	resp, err := g.executeVersion(ctx, "POST", "oauth/access_tokens", request, apiVersion6)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.oauthCredentials()
	}
	return nil, &invalidResponseError{resp}
}

func (g *OAuthGateway) RevokeAccessToken(ctx context.Context, request *OAuthRevokeAccessTokenRequest) (*OAuthRevokeToken, error) {
	resp, err := g.executeVersion(ctx, "POST", "oauth/revoke_access_token", request, apiVersion6)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.revokeAccessTokenResponse()
	}
	return nil, &invalidResponseError{resp}
}
