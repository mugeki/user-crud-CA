package request

import "user-crud-ca/business/users"

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	ClientID int    `json:"client_id"`
}

func (req *User) ToDomain() *users.Domain {
	return &users.Domain{
		ID:       req.ID,
		Name:     req.Name,
		Age:      req.Age,
		Sex:      req.Sex,
		ClientID: req.ClientID,
	}
}