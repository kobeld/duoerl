package services

import (
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
)

func GetUser(userId string) (apiUser *duoerlapi.User, err error) {
	userOId, err := utils.ToObjectId(userId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	user, err := users.FindById(userOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiUser = toApiUser(user)

	return
}

func GetUsers() (apiUsers []*duoerlapi.User, err error) {
	dbUsers, err := users.FindAll(nil)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiUsers = toApiUsers(dbUsers)

	return
}

func toApiUsers(dbUsers []*users.User) (apiUsers []*duoerlapi.User) {
	for _, dbUser := range dbUsers {
		apiUsers = append(apiUsers, toApiUser(dbUser))
	}
	return
}

func toApiUser(user *users.User) *duoerlapi.User {
	apiUser := new(duoerlapi.User)

	if user != nil {
		apiUser = &duoerlapi.User{
			Id:      user.Id.Hex(),
			Link:    user.Link(),
			Name:    user.Name,
			Email:   user.Email,
			Profile: toApiProfile(user.Profile),
		}
	}

	return apiUser
}

func toApiProfile(profile users.Profile) *duoerlapi.Profile {
	apiProfile := &duoerlapi.Profile{
		Gender:          profile.GenderText(),
		Description:     profile.Description,
		Location:        profile.Location,
		Birthday:        profile.BirthdayText(),
		SkinTexture:     profile.SkinTexture,
		SkinTextureText: global.SkinTextureOptions[profile.SkinTexture],
		HairTexture:     profile.HairTexture,
		HairTextureText: global.HairTextureOptions[profile.HairTexture],
	}

	return apiProfile
}
