package efficacies

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/govalidations"
)

type EfficacyGateKeeper struct {
	govalidations.GateKeeper
}

func (this *Efficacy) ValidateCreation() *govalidations.Validated {
	gk := &EfficacyGateKeeper{}
	gk.AddNameValidator()
	gk.AddUniqueValidator()
	return gk.Validate(this)
}

func (this *EfficacyGateKeeper) AddNameValidator() {
	this.Add(govalidations.Presence(func(object interface{}) interface{} {
		return object.(*Efficacy).Name
	}, "Name", global.EFFICACY_02))

	return
}

func (this *EfficacyGateKeeper) AddUniqueValidator() {
	this.Add(govalidations.Custom(func(object interface{}) bool {
		efficacy := object.(*Efficacy)
		c, _ := FindUniqueEfficacy(efficacy.Name, efficacy.ParentId)
		return c == nil
	}, "Name", global.EFFICACY_01))

	return
}
