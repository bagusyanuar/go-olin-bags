package repositories

type AuthRepository interface {
	SignIn() error
}
