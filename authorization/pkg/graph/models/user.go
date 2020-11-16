package models

import (
	"authorization/pkg/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	ID        int `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}


func (User) IsEntity() {}

type UserInput struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u User) CreateUser() *gorm.DB {

	hashedPassword, err := HashPassword(u.Password)
	u.Password = hashedPassword
	if err != nil {
		log.Fatal(err)
	}
	return db.Database.Create(&u)
}

func GetAllUsers() []*User {
	var users []*User
	db.Database.Find(&users)
	return users
}
func (u User) GetUserIdByUsername(username string) (int, error) {
	var users []*User
	db.Database.Where("username = ?", username).First(&users)

	user := users[0]
	//if err != nil {
	//	log.Fatal(err)
	//}
	return user.ID, nil
}

func (u *User) Authenticate() bool {
	var users []*User
	 db.Database.Where("username = ?", u.Username).First(&users)
	//if err != nil {
	//	log.Fatal(err)
	//}
	hashedPassword := users[0].Password

	return CheckPasswordHash(u.Password, hashedPassword)
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}