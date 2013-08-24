package services

import (
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/models/articles"
	"github.com/kobeld/duoerl/models/brands"
	"github.com/kobeld/duoerl/models/news"
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerlapi"
	"github.com/theplant/qortex/utils"
	"labix.org/v2/mgo/bson"
)

func NewNews() (newsInput *duoerlapi.NewsInput) {
	newsInput = &duoerlapi.NewsInput{
		Id: bson.NewObjectId().Hex(),
	}
	return
}

func GetNewsInBrand(brandIdHex string) (apiNews []*duoerlapi.News, err error) {
	brandId, err := utils.ToObjectId(brandIdHex)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	newz, err := news.FindSomeByBrandId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	for _, dbNews := range newz {
		apiNews = append(apiNews, toApiNews(dbNews, nil, nil))
	}

	return
}

func CreateNews(input *duoerlapi.NewsInput) (originInput *duoerlapi.NewsInput, err error) {
	originInput = input

	newsId, err := utils.ToObjectId(input.Id)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brandId, err := utils.ToObjectId(input.BrandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	authorId, err := utils.ToObjectId(input.AuthorId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	dbNews := &news.News{
		Id:      newsId,
		BrandId: brandId,
		Article: *articles.NewArticle(input.Title, input.Content, authorId),
	}

	if err = dbNews.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func ShowNews(newsIdHex, userIdHex string) (apiNews *duoerlapi.News, err error) {
	newsId, err := utils.ToObjectId(newsIdHex)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	dbNews, err := news.FindById(newsId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brand, err := brands.FindById(dbNews.BrandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	author, err := users.FindById(dbNews.AuthorId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiNews = toApiNews(dbNews, brand, author)

	return
}

// ----- Private -----

func toApiNews(dbNews *news.News, brand *brands.Brand, author *users.User) *duoerlapi.News {
	apiNews := new(duoerlapi.News)
	if dbNews != nil {
		apiNews = &duoerlapi.News{
			Id:        dbNews.Id.Hex(),
			Title:     dbNews.Title,
			Content:   dbNews.Content,
			Brand:     toApiBrand(brand),
			Author:    toApiUser(author),
			Link:      dbNews.Link(),
			CreatedAt: dbNews.CreatedAt.Format(global.CREATED_AT_LONG),
		}
	}

	return apiNews
}
