package user_cases

type AuthService interface {
	Authenticate(username, password string) (string, error)
	ParseToken(accessToken string) (int, string, error)
}
