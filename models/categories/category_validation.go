package categories

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/govalidations"
)

type CategoryGateKeeper struct {
	govalidations.GateKeeper
}

func (this *Category) ValidateCreation() *govalidations.Validated {
	gk := &CategoryGateKeeper{}
	gk.AddNameAndLevelValidator()
	gk.AddUniqueValidator()
	return gk.Validate(this)
}

func (this *CategoryGateKeeper) AddNameAndLevelValidator() {
	this.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*Category).Name
	}, "Name", global.CATEGORY_02))

	this.Add(govalidations.Custom(func(object interface{}) bool {
		return IsValidLevel(object.(*Category).Level)
	}, "Level", global.CATEGORY_03))

	return
}

func (this *CategoryGateKeeper) AddUniqueValidator() {
	this.Add(govalidations.Custom(func(object interface{}) bool {
		category := object.(*Category)
		c, _ := FindByNameAndLevel(category.Name, category.Level)
		return c == nil
	}, "Name", global.CATEGORY_01))

	return
}
