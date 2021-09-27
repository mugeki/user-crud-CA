package users

import (
	"context"
)

type userService struct {
	userRepository Repository
}

func NewUserService(repo Repository) Service {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) GetByUserID(ctx context.Context, id int) (Domain, error) {
	res, err := service.userRepository.GetByUserID(ctx, id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *userService) GetAll(ctx context.Context) ([]Domain, error) {
	res, err := service.userRepository.GetAll(ctx)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *userService) Store(ctx context.Context, domain *Domain) (Domain, error) {
	res, err := service.userRepository.Store(ctx, domain)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *userService) Update(ctx context.Context, id int, domain *Domain) (Domain, error) {
	res, err := service.userRepository.Update(ctx, id, domain)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *userService) Delete(ctx context.Context, id int) error {
	err := service.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}