package use_cases

type AuthUseCase interface {
	Authenticate(username, password string) (string, error)
	ParseToken(accessToken string) (int, string, error)
}
