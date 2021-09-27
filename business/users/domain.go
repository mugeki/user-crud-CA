package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	ID        int
	Name      string
	Age       int
	Sex       string
	ClientID  int
}

type Service interface{
	GetByUserID(ctx context.Context, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Store(ctx context.Context, data *Domain) (Domain, error)
	Update(ctx context.Context, id int, data *Domain) (Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface{
	GetByUserID(ctx context.Context, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Store(ctx context.Context, data *Domain) (Domain, error)
	Update(ctx context.Context, id int, data *Domain) (Domain, error)
	Delete(ctx context.Context, id int) error
}