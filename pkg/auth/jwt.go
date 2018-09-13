package auth

import (
	"context"
	"errors"
	"net/http"

	log "github.com/mtbarta/monocorpus/pkg/logging"
)

func TeamAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logger.Log(r.Context().Value("user").(interface{}))
		claims, err := getClaims(r)

		if err != nil {
			log.Logger.Fatalf("claimsError", err.Error(), "request", r.Context())
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		if claims.Email == "" {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}

		newRequest := r.WithContext(context.WithValue(r.Context(), "email", claims.Email))

		next.ServeHTTP(w, newRequest)
	})
}

// func QueryWithinTeamBounds(query notes.Query, r *http.Request) bool {
// 	team := query.Team
// 	claims, err := getClaims(r)

// 	if err != nil {
// 		return false
// 	}

// 	if team == claims. {
// 		return true
// 	}

// 	return false
// }

func getClaims(r *http.Request) (*Claims, error) {
	// user := context.Get(r, "user")
	user := r.Context().Value("user")

	if user == nil {
		log.Logger.Fatalf("err", "no user token set", "user", user)
		return nil, errors.New("no user token set")
	}
	claims := user.(Claims)

	return &claims, nil
}
