package efficacies

import (
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
)

const (
	EFFICACIES = "efficacies"
)

func (this *Efficacy) Save() error {
	return mgodb.Save(EFFICACIES, this)
}

func FindUniqueEfficacy(name string, parentId bson.ObjectId) (r *Efficacy, err error) {
	query := bson.M{"name": name, "parentid": parentId}
	return FindOne(query)
}

func FindAll(query bson.M) (efficacies []*Efficacy, err error) {
	err = mgodb.FindAll(EFFICACIES, query, &efficacies)
	return
}

func FindOne(query bson.M) (r *Efficacy, err error) {
	err = mgodb.FindOne(EFFICACIES, query, &r)
	return
}
