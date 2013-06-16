package accounts

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/govalidations"
	"regexp"
)

type AccountGateKeeper struct {
	govalidations.GateKeeper
}

func (this *Account) ValidateSignupForm() *govalidations.Validated {
	mgk := &AccountGateKeeper{}
	mgk.AddEmailValidator()
	mgk.AddNameValidator()
	mgk.AddPwdPresenceValidator()
	mgk.AddPwdLengthValidator()
	mgk.AddConfirmPasswordValidator()
	return mgk.Validate(this)
}

func (this *Account) ValidateEmailExist() *govalidations.Validated {
	mgk := &AccountGateKeeper{}
	mgk.AddEmailExistValidator()
	return mgk.Validate(this)

}

func (this *Account) ValidateLoginForm() *govalidations.Validated {
	mgk := &AccountGateKeeper{}
	mgk.AddEmailValidator()
	mgk.AddPwdPresenceValidator()
	return mgk.Validate(this)
}

func (this *Account) ValidateLoginAccount() *govalidations.Validated {
	mgk := &AccountGateKeeper{}
	mgk.AddPasswordMatchValidator(this)
	return mgk.Validate(this)
}

func (agk *AccountGateKeeper) AddPasswordMatchValidator(a *Account) {
	email, password := a.Email, a.Password
	account, _ := FindByEmail(email)
	agk.Add(govalidations.Custom(func(object interface{}) bool {
		if account == nil {
			return false
		}
		return true
	}, "Password", global.ACCOUNT_01))

	agk.Add(govalidations.Custom(func(object interface{}) bool {
		if account != nil && !account.IsPwdMatch(password) {
			return false
		}
		return true
	}, "Password", global.ACCOUNT_01))

	a = account
	return
}

func (agk *AccountGateKeeper) AddEmailExistValidator() {
	agk.Add(govalidations.Custom(func(object interface{}) bool {
		email := object.(*Account).Email
		if account, _ := FindByEmail(email); account != nil {
			return false
		}
		return true
	}, "Email", global.ACCOUNT_02))

	return
}

func (agk *AccountGateKeeper) AddEmailValidator() {
	agk.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*Account).Email
	}, "Email", global.ACCOUNT_03))

	agk.Add(govalidations.Regexp(func(object interface{}) interface{} {
		return object.(*Account).Email
	}, regexp.MustCompile(global.EMAIL_REGEXP), "Email", global.ACCOUNT_04))

	return
}

func (agk *AccountGateKeeper) AddNameValidator() {
	agk.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*Account).Name
	}, "Name", global.ACCOUNT_05))

	agk.Add(govalidations.Limitation(func(object interface{}) interface{} {
		return object.(*Account).Name
	}, 0, 20, "Name", global.ACCOUNT_06))

	return
}

func (agk *AccountGateKeeper) AddPwdLengthValidator() {
	agk.Add(govalidations.Prohibition(func(object interface{}) interface{} {
		return object.(*Account).Password
	}, 1, 5, "Password", global.ACCOUNT_07))

	return
}

func (agk *AccountGateKeeper) AddPwdPresenceValidator() {
	agk.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*Account).Password
	}, "Password", global.ACCOUNT_08))

	return
}

func (agk *AccountGateKeeper) AddConfirmPasswordValidator() {
	agk.Add(govalidations.Custom(func(object interface{}) bool {
		p := object.(*Account)
		if p.Password != p.ConfirmPassword {
			return false
		}
		return true
	}, "ConfirmPassword", global.ACCOUNT_09))
	return
}
