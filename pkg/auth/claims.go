package auth

type Claims struct {
	Email              string
	Name               string
	Family_name        string
	Given_name         string
	Preferred_username string
	Groups             []string
	Roles              []string
	Acr                string
	AllowedOrigins     []string `json:allowed-origins`
	Authtime           string   `json:auth_time`
	Exp                int
	iat                int
	iss                string
	jti                string
	nbf                int
	nonce              string
	Realm_Access       map[string][]string
	Session_State      string
	Sub                string
	Type               string
}

var UserClaims struct {
	Email              string
	Name               string
	Family_name        string
	Given_name         string
	Preferred_username string
	Groups             []string
	Roles              []string
	Acr                string
	AllowedOrigins     []string `json:allowed-origins`
	Authtime           string   `json:auth_time`
	Exp                int
	iat                int
	iss                string
	jti                string
	nbf                int
	nonce              string
	Realm_Access       map[string][]string
	Session_State      string
	Sub                string
	Type               string
}
