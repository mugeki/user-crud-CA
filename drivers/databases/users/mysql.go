package users

import (
	"context"
	"user-crud-ca/business/users"

	"gorm.io/gorm"
)

type mysqlUsersRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository{
	return &mysqlUsersRepository{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlUsersRepository) GetByUserID(ctx context.Context, id int) (users.Domain, error){
	rec := User{}
	err := mysqlRepo.Conn.First(&rec, id).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlUsersRepository) GetAll(ctx context.Context) ([]users.Domain, error){
	rec := []User{}
	err := mysqlRepo.Conn.Find(&rec).Error
	if err != nil {
		return []users.Domain{}, err
	}

	var domainUsers []users.Domain
	for _, val := range rec{
		domainUsers = append(domainUsers,val.toDomain())
	}
	return domainUsers, nil
}

func (mysqlRepo *mysqlUsersRepository) Store(ctx context.Context, data *users.Domain) (users.Domain, error){
	rec := fromDomain(*data)
	err := mysqlRepo.Conn.Save(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlUsersRepository) Update(ctx context.Context, id int, data *users.Domain) (users.Domain, error){
	rec := User{}
	newData := fromDomain(*data)
	err := mysqlRepo.Conn.First(&rec, id).Save(&newData).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlUsersRepository) Delete(ctx context.Context, id int) error{
	rec := User{}
	err := mysqlRepo.Conn.First(&rec, id).Delete(&rec).Error
	if err != nil {
		return err
	}
	return nil
}