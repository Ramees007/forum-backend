package usecase

import (
	"github.com/rameesThattarath/qaForum/entity"
	"github.com/rameesThattarath/qaForum/repository"
)

// RegisterUser - RegisterUser
type RegisterUser interface {
	RegisterUser(entity.User) (uint, error)
}

//RegisterUserImpl - RegisterUserImpl
type RegisterUserImpl struct {
	Repo repository.UserRepo
}

//RegisterUser - RegisterUser
func (uc *RegisterUserImpl) RegisterUser(u entity.User) (uint, error) {
	id, err := uc.Repo.Register(u)
	if err != nil {
		return 0, err
	}

	return id, nil

}
