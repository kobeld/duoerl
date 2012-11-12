package accounts

import (
	"github.com/sunfmin/govalidations"
	"github.com/theplant/qortex/configs"
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
	account, _ := FindByEmail(a.Email)
	errorText := "Account and password do not match!"
	agk.Add(govalidations.Custom(func(object interface{}) bool {
		if account == nil {
			return false
		}
		return true
	}, "Password", errorText))

	agk.Add(govalidations.Custom(func(object interface{}) bool {
		if account != nil && !account.IsPwdMatch(a.Password) {
			return false
		}
		return true
	}, "Password", errorText))

	return
}

func (agk *AccountGateKeeper) AddEmailExistValidator() {
	agk.Add(govalidations.Custom(func(object interface{}) bool {
		email := object.(*Account).Email
		if account, _ := FindByEmail(email); account != nil {
			return false
		}
		return true
	}, "Email", "Email already be taken!"))

	return
}

func (agk *AccountGateKeeper) AddEmailValidator() {
	agk.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*Account).Email
	}, "Email", "Email can't be blank!"))

	agk.Add(govalidations.Regexp(func(object interface{}) interface{} {
		return object.(*Account).Email
	}, regexp.MustCompile(configs.EMAIL_REGEXP), "Email", "Format is error!"))

	return
}

func (agk *AccountGateKeeper) AddNameValidator() {
	agk.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*Account).Name
	}, "Name", "Name can't be blank"))

	agk.Add(govalidations.Limitation(func(object interface{}) interface{} {
		return object.(*Account).Name
	}, 0, 20, "Name", "The length of First Name must less than 20"))

	return
}

func (agk *AccountGateKeeper) AddPwdLengthValidator() {
	agk.Add(govalidations.Prohibition(func(object interface{}) interface{} {
		return object.(*Account).Password
	}, 1, 5, "Password", "Password is too short"))

	return
}

func (agk *AccountGateKeeper) AddPwdPresenceValidator() {
	agk.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*Account).Password
	}, "Password", "Password can't be blank"))

	return
}

func (agk *AccountGateKeeper) AddConfirmPasswordValidator() {
	agk.Add(govalidations.Custom(func(object interface{}) bool {
		p := object.(*Account)
		if p.Password != p.ConfirmPassword {
			return false
		}
		return true
	}, "ConfirmPassword", "Password does not match"))
	return
}
