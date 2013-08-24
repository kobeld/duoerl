package categories

import (
	"github.com/kobeld/duoerl/models/attributes"
	"github.com/kobeld/duoerlapi"
	"labix.org/v2/mgo/bson"
)

const (
	LEVEL_ONE string = "Lev1"
	LEVEL_TWO        = "Lev2"
)

type Category struct {
	Id                   bson.ObjectId `bson:"_id"`
	attributes.Attribute `,inline`
	Level                string
}

func (this *Category) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func NewCategory(input *duoerlapi.CategoryInput) *Category {
	return &Category{
		Attribute: *attributes.NewCategoryAttr(input.Name, input.ParentId),
		Level:     input.Level,
	}
}

func IsValidLevel(level string) bool {
	return level == LEVEL_ONE || level == LEVEL_TWO
}
