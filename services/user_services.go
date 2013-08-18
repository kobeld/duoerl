package services

import (
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
	"time"
)

func UpdateProfile(userInput *duoerlapi.UserInput) (err error) {
	user, err := users.FetchByIdHex(userInput.Id)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	user.AvatarUrl = userInput.Avatar
	user.Profile.Gender = userInput.Profile.Gender
	user.Profile.Location = userInput.Profile.Location
	user.Profile.Description = userInput.Profile.Description
	user.Profile.HairTexture = userInput.Profile.HairTexture
	user.Profile.SkinTexture = userInput.Profile.SkinTexture
	user.Profile.Birthday, _ = time.Parse(global.DATE_BIRTHDAY, userInput.Profile.Birthday)

	if err = user.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func GetUser(userId string) (apiUser *duoerlapi.User, err error) {

	user, err := users.FetchByIdHex(userId)
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
			Avatar:  user.Avatar(),
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
		Birthday:        profile.Birthday.Format(global.DATE_BIRTHDAY),
		SkinTexture:     profile.SkinTexture,
		SkinTextureText: global.SkinTextureOptions[profile.SkinTexture],
		HairTexture:     profile.HairTexture,
		HairTextureText: global.HairTextureOptions[profile.HairTexture],
	}

	return apiProfile
}
