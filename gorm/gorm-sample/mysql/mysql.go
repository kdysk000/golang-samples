package mysql

import (
	"gorm-sample/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DbInit() {
	var err error
	dsn := "root:secret@tcp(127.0.0.1:3306)/sample_db?charset=utf8mb4&parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func AutoMigrate(models ...interface{}) {
	db.AutoMigrate(models...,)
}

func AddUser(user model.User) error {
	_user := model.User{
		UserId:  user.UserId,
		Name:    user.Name,
		Age:     user.Age,
		Address: user.Address,
		Tasks:   []model.Task{
			{
				TaskId: "task-0001",
				Name:   "task1",
			},
			{
				TaskId: "task-0002",
				Name:   "task2",
			},
		},
	}
	ret := db.Create(&_user)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	return ret.Error
}

func GetUser() ([]model.User, error) {
	users := []model.User{}
	ret := db.Find(&users)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	return users, nil
}

func UpdateUser(user model.User) error {
	result := db.Model(&model.User{}).
		Where("user_id = ?", user.UserId).
		Updates(model.User{Name: user.Name, Age: user.Age, Address: user.Address})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	users := []model.User{}
	ret := db.Find(&users)
	if ret.Error != nil {
		log.Fatal(ret.Error)
		return ret.Error
	}
	return nil
}

func DeleteUser(userid string) error {
	user := model.User{}
	ret := db.Where("user_id = ?", userid).First(&user)
	if ret.Error != nil {
		log.Fatal(ret.Error)
		return ret.Error
	}
	ret = db.Delete(&user)
	if ret.Error != nil {
		log.Fatal(ret.Error)
		return ret.Error
	}
	return nil
}
