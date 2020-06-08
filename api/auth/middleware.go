package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// we need a type for context key
type ctxKey string

var (
	// TokenClaimsKey is the key in the request
	// context object for auth token claims
	TokenClaimsKey = ctxKey("claims")
)

// Wrap wraps an HTTP handler function
// and populates the access token claims object in the req ctx
func (a *Authenticator) Wrap(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHdr := r.Header.Get("Authorization")
		jwt := strings.TrimPrefix(authHdr, "Bearer ")

		if authHdr == jwt {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "no bearer token in header")
			return
		}

		verifiedClaims, err := a.ValidateJWT(jwt)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "invalid access token")
			return
		}

		// run handler with token in context
		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), TokenClaimsKey, verifiedClaims)))
	})
}

// GetClaims returns the claims in a context object
func GetClaims(r *http.Request) *CustomClaims {
	return r.Context().Value(TokenClaimsKey).(*CustomClaims)
}
