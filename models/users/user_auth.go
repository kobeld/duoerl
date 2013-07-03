package users

import (
	"code.google.com/p/go.crypto/bcrypt"
	"time"
)

func LoginWith(email string, pwd string) (user *User) {
	a, _ := FindByEmail(email)
	if a == nil {
		return
	}
	if a.IsPwdMatch(pwd) {
		user = a
	}
	return
}

func (this *User) IsPwdMatch(pwd string) bool {
	if bcrypt.CompareHashAndPassword([]byte(this.Password), []byte(pwd)) != nil {
		return false
	}
	return true
}

func (this *User) Signup() (err error) {
	this.CreatedAt = time.Now()
	this.encryptPwd()
	return this.Save()
}

func (this *User) encryptPwd() {
	hp, _ := bcrypt.GenerateFromPassword([]byte(this.Password), 0)
	this.Password = string(hp)
}
