package firebase

import (
	"context"

	f "firebase.google.com/go"
)

type IAuth interface {
	getUid(string) (string, error)
	Authenticate(string) error
}

type Auth struct {
	fa *f.App
}

func NewAuth(afa *f.App) (*Auth, error) {
	return &Auth{fa: afa},nil
}

func (a *Auth) GetUid(it string) (string, error) {
	cli, err := a.fa.Auth(context.Background())
	if err != nil {
		return "", err
	}

	token, err := cli.VerifyIDToken(context.TODO(), it)
	if err != nil {
		return "", err
	}

	return token.UID, nil
}
