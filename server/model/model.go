package model

import (
	"fmt"
	"github.com/zqhhh/gomall/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"runtime"
)

type Model struct {
	ID uint `json:"id" gorm:"primary_key"`
}

var (
	db *gorm.DB
	DB *gorm.DB
)

func connect() {
	var err error
	config := app.Config.DB
	db, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Name,
		))
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.LogMode(true)
	if runtime.GOOS == "windows" {
		// windows不支持日志字体颜色
		db.SetLogger(myLogger{})
	}
	DB = db
}

func migrate() {
	db.AutoMigrate(
		&Category{},
		&AD{},
		&Channel{},
		&Brand{},
		&Goods{},
		&Topic{},
		&Gallery{},
		&Issue{},
		&User{},
		&Cart{},
		&AttributeKey{},
		&AttributeValue{},
		&SpecificationKey{},
		&SpecificationValue{},
		&Product{},
		&Order{},
		&OrderGoods{},
		&UserAddress{},
		&Region{},
	)
}

func InitDB() {
	connect()
	migrate()
}
