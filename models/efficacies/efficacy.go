package efficacies

import (
	"github.com/kobeld/duoerl/models/attributes"
	"github.com/kobeld/duoerlapi"
	"labix.org/v2/mgo/bson"
)

type Efficacy struct {
	attributes.Attribute `,inline`
}

func (this *Efficacy) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func NewEfficacy(input *duoerlapi.EfficacyInput) *Efficacy {
	return &Efficacy{
		Attribute: *attributes.NewEfficacyAttr(input.Name, input.ParentId),
	}
}
