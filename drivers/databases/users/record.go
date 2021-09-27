package users

import (
	"user-crud-ca/business/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Sex      string
	ClientID int
}

func (rec *User) toDomain() users.Domain{
	return users.Domain{
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		CreatedAt: rec.CreatedAt,
		ID: int(rec.ID),
		Name: rec.Name,
		Age: rec.Age,
		Sex: rec.Sex,
		ClientID: rec.ClientID,
	}
}

func fromDomain(domain users.Domain) *User{
	return &User{
		gorm.Model{
			UpdatedAt: domain.UpdatedAt,
			DeletedAt: domain.DeletedAt,
			CreatedAt: domain.CreatedAt,
			ID: uint(domain.ID),
		},
		domain.Name,
		domain.Age,
		domain.Sex,
		domain.ClientID,
	}
}