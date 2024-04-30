package auth

type Authenticator interface {
	Authenticate(*AuthenticationRequest) any
}
