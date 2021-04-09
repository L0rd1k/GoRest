package models

import (
	"bookmarks-api/utils"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Права доступа JWT
type Token struct {
	UserId uint
	jwt.StandardClaims
}

// Учетная запись пользователя
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token" ;sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}
	if len(account.Password) < 6 {
		return utils.Message(false, "Password required! Less than 6!"), false
	}
	temp := &Account{}
	err := GetDataBase().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "Email address already in use by another user."), false
	}
	return utils.Message(false, "Requirement passed"), true
}

func (account *Account) Create() map[string]interface{} {
	if resp, ok := account.Validate(); !ok {
		return resp
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	GetDataBase().Create(account)
	if account.ID <= 0 {
		return utils.Message(false, "Failed to create account, connection error.")
	}
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = ""
	response := utils.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := GetDataBase().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email address not found")
		}
		return utils.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Пароль не совпадает!!
		return utils.Message(false, "Invalid login credentials. Please try again")
	}
	// WORKS! LOGIN TO SYSTEM
	account.Password = ""

	// CREATE JWT TOKEN
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString // Сохраните токен в ответе

	resp := utils.Message(true, "Logged In")
	resp["account"] = account
	return resp

}

func GetUser(u uint) *Account {
	acc := &Account{}
	GetDataBase().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil
	}
	acc.Password = ""
	return acc
}
