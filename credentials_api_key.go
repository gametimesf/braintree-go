package braintree

import "encoding/base64"

type apiKey struct {
	env          Environment
	graphqlEnv   Environment
	merchantID   string
	publicKey    string
	privateKey   string
	clientSecret string
	clientId     string
}

func newAPIKey(env, graphqlEnv Environment, merchantID, publicKey, privateKey string) credentials {
	return apiKey{
		env:        env,
		graphqlEnv: graphqlEnv,
		merchantID: merchantID,
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func newAPIKeyWithCredentials(env, graphqlEnv Environment, merchantID, publicKey, privateKey, clientId, clientSecret string) credentials {
	return apiKey{
		env:        env,
		graphqlEnv: graphqlEnv,
		merchantID: merchantID,
		publicKey:  publicKey,
		privateKey: privateKey,

		// external creds for limited access to our partner external braintree.
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (k apiKey) Environment() Environment {
	return k.env
}

func (k apiKey) GraphQLEnvironment() Environment {
	return k.graphqlEnv
}

func (k apiKey) MerchantID() string {
	return k.merchantID
}

func (k apiKey) AuthorizationHeader() string {
	auth := k.publicKey + ":" + k.privateKey
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

// AuthorizationHeaderV2 uses the client id and client secret for the Authorization header
func (k apiKey) AuthorizationHeaderWithClientCreds() string {
	auth := k.clientId + ":" + k.clientSecret
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
