package drivers

import (
	userDomain "user-crud-ca/business/users"
	userDB "user-crud-ca/drivers/databases/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}