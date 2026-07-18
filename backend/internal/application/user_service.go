package application

import (
	"fmt"

	"github.com/AoiSatonaka/todo-app/backend/internal/constant"
	"github.com/AoiSatonaka/todo-app/backend/internal/domain/user"
)

type IUserService interface {
	Authorize(string) (*user.User, error)
}

type UserService struct {
	ur user.IUserRepository
}

var _ IUserService = &UserService{}

func NewUserService(aur user.IUserRepository) *UserService {
	return &UserService{ur: aur}
}

func (us *UserService) Authorize(idt string) (*user.User, error) {
	if idt == "" {
		return nil, fmt.Errorf(constant.ERR_AUTH_IDTOKEN_NOT_SET)
	}
	user, err := us.ur.Get(idt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
