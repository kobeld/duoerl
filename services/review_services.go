package services

import (
	"github.com/kobeld/duoerl/models/products"
	"github.com/kobeld/duoerl/models/reviews"
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
	"labix.org/v2/mgo/bson"
)

/*
ReviewInput:
	Id        string
	ProductId string
	AuthorId  string
	Content   string
*/

func NewReview() *duoerlapi.ReviewInput {
	return &duoerlapi.ReviewInput{Id: bson.NewObjectId().Hex()}
}

// Todo: Validation Needed
func CreateReview(input *duoerlapi.ReviewInput) (originInput *duoerlapi.ReviewInput, err error) {
	originInput = input

	oId, err := utils.ToObjectId(input.Id)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	productOId, err := utils.ToObjectId(input.ProductId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	// Check if the product exists
	product, err := products.FindById(productOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	authorOId, err := utils.ToObjectId(input.AuthorId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	review := &reviews.Review{
		Id:          oId,
		AuthorId:    authorOId,
		ProductId:   productOId,
		BrandId:     product.BrandId,
		Content:     input.Content,
		Rating:      input.Rating,
		EfficacyIds: utils.TurnPlainIdsToObjectIds(input.EfficacyIds),
	}

	if err = review.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func ShowReviewsInProduct(productId string) (apiReviews []*duoerlapi.Review, err error) {

	productOId, err := utils.ToObjectId(productId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	dbReviews, err := reviews.FindSomeByProductId(productOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	authorIds, productIds := reviews.CollectAuthorAndProductIds(dbReviews)

	dbAuthors, err := users.FindByIds(authorIds)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	dbProducts, err := products.FindByIds(productIds)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	productMap := products.BuildProductMap(dbProducts)
	authorMap := users.BuildUserMap(dbAuthors)

	apiReviews = toApiReviews(dbReviews, productMap, authorMap)

	return
}

// ------------- Private  ---------------

func toApiReviews(dbReviews []*reviews.Review, productMap map[bson.ObjectId]*products.Product,
	authorMap map[bson.ObjectId]*users.User) (apiReviews []*duoerlapi.Review) {

	for _, dbReview := range dbReviews {
		dbProduct := productMap[dbReview.ProductId]
		dbAuthor := authorMap[dbReview.AuthorId]
		apiReviews = append(apiReviews, toApiReview(dbReview, dbProduct, dbAuthor))
	}

	return
}

func toApiReview(dbReview *reviews.Review, dbProduct *products.Product, dbAuthor *users.User) *duoerlapi.Review {
	apiReview := new(duoerlapi.Review)
	if dbReview != nil {
		efficacyIds := utils.TurnObjectIdToPlainIds(dbReview.EfficacyIds)
		apiReview = &duoerlapi.Review{
			Id:         dbReview.Id.Hex(),
			Content:    dbReview.Content,
			Product:    toApiProduct(dbProduct, nil, dbAuthor),
			Author:     toApiUser(dbAuthor),
			Rating:     dbReview.Rating,
			Efficacies: GetEfficaciesByIds(efficacyIds),
		}
	}
	return apiReview
}
