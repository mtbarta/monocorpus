package gateway

import (
	"context"
	"errors"
	"net/http"

	notes "github.com/mtbarta/monocorpus/pkg/notes/types"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

func OnError(w http.ResponseWriter, r *http.Request, err string) {
	logger.Log("error", err)
	http.Error(w, err, http.StatusExpectationFailed)
}

func NewJWTMiddleware(pemBytes []byte) *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return jwt.ParseRSAPublicKeyFromPEM(pemBytes)
		},
		Debug:        false,
		ErrorHandler: OnError,
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodRS256,
	})
}

func TeamAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logger.Log(r.Context().Value("user").(interface{}))
		claims, err := getClaims(r)

		if err != nil {
			logger.Log("claimsError", err.Error(), "request", r.Context())
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		if claims["email"] == nil {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}

		newRequest := r.WithContext(context.WithValue(r.Context(), "email", claims["email"]))

		next.ServeHTTP(w, newRequest)
	})
}

func QueryWithinTeamBounds(query notes.Query, r *http.Request) bool {
	team := query.Team
	claims, err := getClaims(r)

	if err != nil {
		return false
	}

	if team == claims["team"] {
		return true
	}

	return false
}

func getClaims(r *http.Request) (jwt.MapClaims, error) {
	// user := context.Get(r, "user")
	user := r.Context().Value("user")

	if user == nil {
		logger.Log("err", "no user token set", "user", user)
		return nil, errors.New("no user token set")
	}
	claims := user.(*jwt.Token).Claims.(jwt.MapClaims)

	return claims, nil
}
