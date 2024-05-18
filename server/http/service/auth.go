package service

type Authenticator interface {
	Authenticate(*AuthenticationRequest) any
}
