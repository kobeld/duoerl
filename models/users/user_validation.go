package users

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/govalidations"
	"regexp"
)

type UserGateKeeper struct {
	govalidations.GateKeeper
}

func (this *User) ValidateSignupForm() *govalidations.Validated {
	mgk := &UserGateKeeper{}
	mgk.AddEmailValidator()
	mgk.AddNameValidator()
	mgk.AddPwdPresenceValidator()
	mgk.AddPwdLengthValidator()
	mgk.AddConfirmPasswordValidator()
	return mgk.Validate(this)
}

func (this *User) ValidateEmailExist() *govalidations.Validated {
	mgk := &UserGateKeeper{}
	mgk.AddEmailExistValidator()
	return mgk.Validate(this)

}

func (this *User) ValidateLoginForm() *govalidations.Validated {
	mgk := &UserGateKeeper{}
	mgk.AddEmailValidator()
	mgk.AddPwdPresenceValidator()
	return mgk.Validate(this)
}

func (this *User) ValidateLoginUser() *govalidations.Validated {
	mgk := &UserGateKeeper{}
	mgk.AddPasswordMatchValidator(this)
	return mgk.Validate(this)
}

func (agk *UserGateKeeper) AddPasswordMatchValidator(user *User) {
	email, password := user.Email, user.Password
	user, _ = FindByEmail(email)
	agk.Add(govalidations.Custom(func(object interface{}) bool {
		if user == nil {
			return false
		}
		return true
	}, "Password", global.USER_01))

	agk.Add(govalidations.Custom(func(object interface{}) bool {
		if user != nil && !user.IsPwdMatch(password) {
			return false
		}
		return true
	}, "Password", global.USER_01))

	return
}

func (agk *UserGateKeeper) AddEmailExistValidator() {
	agk.Add(govalidations.Custom(func(object interface{}) bool {
		email := object.(*User).Email
		if user, _ := FindByEmail(email); user != nil {
			return false
		}
		return true
	}, "Email", global.USER_02))

	return
}

func (agk *UserGateKeeper) AddEmailValidator() {
	agk.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*User).Email
	}, "Email", global.USER_03))

	agk.Add(govalidations.Regexp(func(object interface{}) interface{} {
		return object.(*User).Email
	}, regexp.MustCompile(global.EMAIL_REGEXP), "Email", global.USER_04))

	return
}

func (agk *UserGateKeeper) AddNameValidator() {
	agk.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*User).Name
	}, "Name", global.USER_05))

	agk.Add(govalidations.Limitation(func(object interface{}) interface{} {
		return object.(*User).Name
	}, 0, 20, "Name", global.USER_06))

	return
}

func (agk *UserGateKeeper) AddPwdLengthValidator() {
	agk.Add(govalidations.Prohibition(func(object interface{}) interface{} {
		return object.(*User).Password
	}, 1, 5, "Password", global.USER_07))

	return
}

func (agk *UserGateKeeper) AddPwdPresenceValidator() {
	agk.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*User).Password
	}, "Password", global.USER_08))

	return
}

func (agk *UserGateKeeper) AddConfirmPasswordValidator() {
	agk.Add(govalidations.Custom(func(object interface{}) bool {
		p := object.(*User)
		if p.Password != p.ConfirmPassword {
			return false
		}
		return true
	}, "ConfirmPassword", global.USER_09))
	return
}
