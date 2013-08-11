package categories

import (
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
)

const (
	CATEGORIES = "categories"
)

func (this *Category) Save() error {
	return mgodb.Save(CATEGORIES, this)
}

func FindByNameAndLevel(name, level string) (r *Category, err error) {
	query := bson.M{"name": name, "level": level}
	return FindOne(query)
}

func FindAll(query bson.M) (categories []*Category, err error) {
	err = mgodb.FindAll(CATEGORIES, query, &categories)
	return
}

func FindOne(query bson.M) (r *Category, err error) {
	err = mgodb.FindOne(CATEGORIES, query, &r)
	return
}
