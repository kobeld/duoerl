package attributes

import (
	"github.com/kobeld/duoerl/utils"
	"labix.org/v2/mgo/bson"
)

const (
	TYPE_CATEGORY = "attr_01"
)

type Attribute struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	AType    string
	ParentId bson.ObjectId `bson:",omitempty" json:",omitempty"`
}

func (this *Attribute) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func NewCategoryAttr(name, parentIdHex string) (attr *Attribute) {
	parentId, _ := utils.ToObjectId(parentIdHex)
	attr = &Attribute{
		Id:       bson.NewObjectId(),
		Name:     name,
		AType:    TYPE_CATEGORY,
		ParentId: parentId,
	}

	return
}
