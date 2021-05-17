package model

type User struct {
	Model
	Username      string `gorm:"UNIQUE"`
	Password      string `gorm:"type:varchar(32)"`
	Birthday      int64
	RegisterTime  int64
	LastLoginTime int64
	LastLoginIP   string
	Nickname      string
	Mobile        string
	RegisterIP    string
	Avatar        string
}

func (u *User) Insert() error {
	return db.Create(u).Error
}

func (u *User) Query(selectFields []string, where map[string]interface{}) error {
	return db.Model(User{}).Select(selectFields).Where(where).First(u).Error
}
