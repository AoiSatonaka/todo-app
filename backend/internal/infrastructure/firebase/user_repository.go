package firebase

import (
	"context"

	f "firebase.google.com/go"
	"github.com/AoiSatonaka/todo-app/backend/internal/domain/user"
)

type UserRepository struct {
	fa *f.App
}

var _ user.IUserRepository = &UserRepository{}

func NewUserRepository(afa *f.App) *UserRepository {
	return &UserRepository{fa: afa}
}

func (ur *UserRepository) Get(idt string) (*user.User, error) {
	cli, err := ur.fa.Auth(context.Background())
	if err != nil {
		return nil, err
	}
  t,err:=cli.VerifyIDToken(context.Background(), idt)
  if err != nil {
    return nil, err
  }
  
  return &user.User{Id:t.UID}, nil
}
