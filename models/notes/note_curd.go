package notes

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
	"time"
)

const NOTES = "notes"

func (this *Note) Save() error {
	if this.CreatedAt.IsZero() {
		this.CreatedAt = time.Now()
	} else {
		this.UpdatedAt = time.Now()
	}
	return mgodb.Save(NOTES, this)
}

func FindSomeByUserId(userId bson.ObjectId) (r []*Note, err error) {
	if !userId.Valid() {
		err = global.InvalidIdError
		return
	}
	return FindAll(bson.M{"authorid": userId})
}

func FindById(id bson.ObjectId) (r *Note, err error) {
	if !id.Valid() {
		err = global.InvalidIdError
		return
	}
	return FindOne(bson.M{"_id": id})
}

func FindOne(query bson.M) (note *Note, err error) {
	err = mgodb.FindOne(NOTES, query, &note)
	return
}

func FindAll(query bson.M) (r []*Note, err error) {
	err = mgodb.FindAll(NOTES, query, &r)
	return
}
