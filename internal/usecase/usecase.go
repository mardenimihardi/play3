package usecase

import (
	types "play3/internal/model"
	repository "play3/internal/repository"
)

type UsecaseInterface interface {
	GetAllDishes() []types.Dishes
}

type Usecase struct {
	Repository repository.RepositoryInterface
}

func InitUsecase(repository repository.Repository) Usecase {
	return Usecase{
		Repository: repository,
	}
}

func (u Usecase) GetAllDishes() []types.Dishes {
	return u.Repository.GetAllDishes()
}
