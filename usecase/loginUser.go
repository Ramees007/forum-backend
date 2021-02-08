package usecase

import (
	"github.com/rameesThattarath/qaForum/entity"
	"github.com/rameesThattarath/qaForum/repository"
)

// LoginUser - LoginUser
type LoginUser interface {
	LoginUser(entity.User) (uint, *entity.User, error)
}

//LoginUserImpl - LoginUserImpl
type LoginUserImpl struct {
	Repo repository.UserRepo
}

//LoginUser - LoginUser
func (uc *RegisterUserImpl) LoginUser(u entity.User) (uint, *entity.User, error) {
	id, user, err := uc.Repo.Login(u)
	if err != nil {
		return id, nil, err
	}

	return id, user, err

}
