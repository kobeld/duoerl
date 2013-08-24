package news

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
	"time"
)

const NEWS = "news"

func (this *News) Save() error {
	if this.CreatedAt.IsZero() {
		this.CreatedAt = time.Now()
	} else {
		this.UpdatedAt = time.Now()
	}
	return mgodb.Save(NEWS, this)
}

func FindSomeByBrandId(brandId bson.ObjectId) (r []*News, err error) {
	if !brandId.Valid() {
		err = global.InvalidIdError
		return
	}
	return FindAll(bson.M{"brandid": brandId})
}

func FindById(id bson.ObjectId) (r *News, err error) {
	if !id.Valid() {
		err = global.InvalidIdError
		return
	}
	return FindOne(bson.M{"_id": id})
}

func FindOne(query bson.M) (r *News, err error) {
	err = mgodb.FindOne(NEWS, query, &r)
	return
}

func FindAll(query bson.M) (r []*News, err error) {
	err = mgodb.FindAll(NEWS, query, &r)
	return
}
