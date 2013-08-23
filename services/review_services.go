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

func ShowReviewsInProduct(productIdHex string) (apiReviews []*duoerlapi.Review, err error) {

	productId, err := utils.ToObjectId(productIdHex)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	reviewz, err := reviews.FindSomeByProductId(productId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiReviews, err = makeApiReviews(reviewz)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func ShowReviewsInBrand(brandIdHex string) (apiReviews []*duoerlapi.Review, err error) {

	brandId, err := utils.ToObjectId(brandIdHex)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	reviewz, err := reviews.FindSomeByBrandId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiReviews, err = makeApiReviews(reviewz)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

// ------------- Private  ---------------

func makeApiReviews(reviewz []*reviews.Review) (apiReviews []*duoerlapi.Review, err error) {

	authorIds, productIds := reviews.CollectAuthorAndProductIds(reviewz)

	authorz, err := users.FindByIds(authorIds)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	productz, err := products.FindByIds(productIds)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	productMap := products.BuildProductMap(productz)
	authorMap := users.BuildUserMap(authorz)

	apiReviews = toApiReviews(reviewz, productMap, authorMap)

	return
}

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
