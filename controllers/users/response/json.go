package request

import (
	"time"
	"user-crud-ca/business/users"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
	ClientID  int    `json:"client_id"`
}

func FromDomain(domain users.Domain) User {
	return User{
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		ID: domain.ID,
		Name: domain.Name,
		Age: domain.Age,
		Sex: domain.Sex,
		ClientID: domain.ClientID,
	}
}