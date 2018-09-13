package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc"
	log "github.com/mtbarta/monocorpus/pkg/logging"
)

type Token struct {
	Subject string
	claims  []byte
}

func (i *Token) Claims(v interface{}) error {
	if i.claims == nil {
		return errors.New("oidc: claims not set")
	}
	return json.Unmarshal(i.claims, v)
}

// NewJWTMiddleware creates a RemoteKeySet from the jwksURL, and then an oidc.Verifier
// using the issuerURL. This is so that the issuer can be in front of a proxy
// while we retrieve the jwksURL from inside our network.
func NewJWTMiddleware(clientID string, contextProperty string, jwksURL string, issuerURL string, next http.Handler) http.Handler {
	keyset := oidc.NewRemoteKeySet(context.Background(), jwksURL)
	verifier := oidc.NewVerifier(
		issuerURL,
		keyset,
		&oidc.Config{ClientID: clientID},
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := FromAuthHeader(r)
		if err != nil {
			log.Logger.Fatalf("no token set in request")
		}
		token, err := verifier.Verify(r.Context(), tokenString)
		if err != nil {
			log.Logger.Error(err)
			log.Logger.Warn("failed to verify token", err)
			// bail.
			return
		}

		claims := Claims{}
		err = token.Claims(&claims)
		if err != nil {
			log.Logger.Fatalf("could not parse claims", err)
		}
		newRequest := r.WithContext(context.WithValue(r.Context(), contextProperty, claims))
		next.ServeHTTP(w, newRequest)
	})
}

func CreateOIDCProvider(ctx context.Context, issuerURL string) (*oidc.Provider, error) {
	provider, err := oidc.NewProvider(ctx, issuerURL)
	if err != nil {
		log.Logger.Error(err)
		log.Logger.Fatalf("failed to connect to oidc issuer endpoint", "error", err.Error(), "issuerURL", issuerURL)
		return nil, err
	}

	return provider, nil
}

func FromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}
