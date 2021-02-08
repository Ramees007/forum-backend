package usecase

import (
	"github.com/rameesThattarath/qaForum/entity"
	"github.com/rameesThattarath/qaForum/repository"
)

type GetProfile interface {
	GetProfile(uint) (*entity.User, error)
}

type GetProfileImpl struct {
	Repo repository.UserRepo
}

func (p *GetProfileImpl) GetProfile(id uint) (*entity.User, error) {

	u, err := p.Repo.GetProfile(id)

	if err != nil {
		return nil, err
	}
	return u, nil
}
