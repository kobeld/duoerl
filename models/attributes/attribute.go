package attributes

import (
	"github.com/kobeld/duoerl/utils"
	"labix.org/v2/mgo/bson"
)

const (
	TYPE_CATEGORY = "attr_01"
	TYPE_EFFICACY = "attr_02"
)

type Attribute struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	AType    string
	ParentId bson.ObjectId `bson:",omitempty" json:",omitempty"`
}

func newAttribute(name, parentIdHex string) *Attribute {
	parentId, _ := utils.ToObjectId(parentIdHex)
	return &Attribute{
		Id:       bson.NewObjectId(),
		Name:     name,
		AType:    TYPE_CATEGORY,
		ParentId: parentId,
	}
}

func NewCategoryAttr(name, parentIdHex string) *Attribute {
	attr := newAttribute(name, parentIdHex)
	attr.AType = TYPE_CATEGORY

	return attr
}

func NewEfficacyAttr(name, parentIdHex string) *Attribute {
	attr := newAttribute(name, parentIdHex)
	attr.AType = TYPE_EFFICACY

	return attr
}
