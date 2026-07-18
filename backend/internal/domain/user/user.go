package user

type User struct {
	Id string
}

func New(i string) *User {
	return &User{Id: i}
}
