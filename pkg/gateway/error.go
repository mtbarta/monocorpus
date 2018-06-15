package gateway

//Auth error
import "net/http"

//base auth error
type AuthError struct {
	Realm string
}

// Error is an implementation of the Error interface.
func (AuthError) Error() string {
	return http.StatusText(http.StatusUnauthorized)
}

func (AuthError) StatusCode() int {
	return http.StatusUnauthorized
}
