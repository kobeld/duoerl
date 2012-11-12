package accounts

import (
	"github.com/sunfmin/govalidations"
	"github.com/theplant/qortex/configs"
	"regexp"
)

type AccountGateKeeper struct {
	govalidations.GateKeeper
}

func (this *Account) ValidateSignup() govalidations.Errors {
	mgk := &AccountGateKeeper{}
	mgk.AddEmailValidator()
	mgk.AddNameValidator()
	mgk.AddPasswordValidator()
	mgk.AddConfirmPasswordValidator()
	return mgk.Validate(this).Errors
}

func (this *Account) ValidateLogin() govalidations.Errors {
	mgk := &AccountGateKeeper{}
	mgk.AddEmailValidator()
	mgk.AddPasswordValidator()
	return mgk.Validate(this).Errors
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

func (agk *AccountGateKeeper) AddPasswordValidator() {
	agk.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*Account).Password
	}, "Password", "Password can't be blank"))

	agk.Add(govalidations.Prohibition(func(object interface{}) interface{} {
		return object.(*Account).Password
	}, 1, 5, "Password", "Password is too short"))

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
