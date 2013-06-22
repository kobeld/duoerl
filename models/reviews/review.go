package reviews

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Review struct {
	Id        bson.ObjectId `bson:"_id"`
	AuthorId  bson.ObjectId
	ProductId bson.ObjectId
	Content   string
	CreatedAt time.Time
}

func (this *Review) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func CollectAuthorAndProductIds(dbReviews []*Review) (authorIds, productids []bson.ObjectId) {
	for _, dbReview := range dbReviews {
		if dbReview.AuthorId.Valid() {
			authorIds = append(authorIds, dbReview.AuthorId)
		}

		if dbReview.ProductId.Valid() {
			productids = append(productids, dbReview.ProductId)
		}
	}

	return
}
