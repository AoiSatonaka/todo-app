package user

type IUserRepository interface {
	Get(string) (*User, error)
}
