package auth

import (
	"context"

	oidc "github.com/coreos/go-oidc"
)

type claimSource struct {
	Endpoint    string `json:"endpoint"`
	AccessToken string `json:"access_token"`
}

func VerifyToken(certUrl string, issuerURL string) func(string) (*oidc.IDToken, error) {
	ctx := context.Background()
	keySet := oidc.NewRemoteKeySet(ctx, certUrl)
	verifier := oidc.NewVerifier(issuerURL, keySet, &oidc.Config{})

	return func(token string) (*oidc.IDToken, error) {
		idToken, err := verifier.Verify(ctx, token)

		return idToken, err
	}
}
