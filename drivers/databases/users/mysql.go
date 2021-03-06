package users

import (
	"context"
	"myuseek/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Register(ctx context.Context, firstname, lastname, username, email, password, bio, profile_pic, subscription_type string) (users.Domain, error) {
	var userDB Users
	userDB.FirstName = firstname
	userDB.LastName = lastname
	userDB.Email = email
	userDB.Password = password
	userDB.Username = username
	userDB.Bio = bio
	userDB.Profile_pic = profile_pic
	userDB.Subscription_type = subscription_type
	result := rep.Conn.Create(&userDB)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return userDB.ToDomain(), nil
}

func (rep *MysqlUserRepository) Login(ctx context.Context, username, password string) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "username = ?",
		username)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil

}

func (rep *MysqlUserRepository) GetUsers(ctx context.Context, domain users.Domain) ([]users.Domain, error) {
	userlist := []Users{}
	search := "%" + domain.Username + "%"
	result := rep.Conn.Where("username LIKE ? OR first_name LIKE ? OR last_name LIKE ?", search, search, search).Find(&userlist)

	if result.Error != nil {
		userlistdomain := []users.Domain{}
		return userlistdomain, result.Error
	}
	userlistdomain := ToListDomain(userlist)
	return userlistdomain, nil

}

func (rep *MysqlUserRepository) GetUserById(ctx context.Context, domain users.Domain) (users.Domain, error) {
	user := Users{}
	result := rep.Conn.Find(&user, domain.Id)

	if result.Error != nil {
		userdomain := users.Domain{}
		return userdomain, result.Error
	}
	userdomain := user.ToDomain()
	return userdomain, nil

}
