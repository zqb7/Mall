package user

import (
	"errors"
	"github.com/zqhhh/gomall/model"
	"github.com/jinzhu/gorm"
)

type service struct {
	userID uint
}

type Service interface {
	Register(user model.User) error
	Login(username, password string) error
	GetUserID() uint
}

func NewUserService() Service {
	return &service{}
}

func (s *service) Register(user model.User) error {
	return user.Insert()
}

func (s *service) Login(username, password string) error {
	user := model.User{}
	err := user.Query([]string{"id", "username", "password"}, map[string]interface{}{"username": username})
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("用户不存在")
	} else if err == nil && user.Password == password {
		s.userID = user.ID
		return nil
	}
	return err
}
func (s *service) GetUserID() uint {
	return s.userID
}
