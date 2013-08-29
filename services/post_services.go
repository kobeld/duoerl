package services

import (
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/models/posts"
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
	"html/template"
)

func CreatePost(input *duoerlapi.PostInput) (originInput *duoerlapi.PostInput, err error) {
	originInput = input

	// simple validation
	if input.Content == "" {
		err = global.CanNotBeBlankError
		return
	}

	postId, err := utils.ToObjectId(input.Id)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	authorId, err := utils.ToObjectId(input.AuthorId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	post := &posts.Post{
		Id:       postId,
		Content:  input.Content,
		AuthorId: authorId,
	}

	if err = post.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func GetUserPosts(userIdHex string) (apiPosts []*duoerlapi.Post, err error) {

	userId, err := utils.ToObjectId(userIdHex)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	postz, err := posts.FindSomeByAuthorId(userId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	for _, post := range postz {
		apiPosts = append(apiPosts, toApiPost(post, nil))
	}

	return
}

//----- Private -----
func toApiPost(post *posts.Post, author *users.User) *duoerlapi.Post {
	apiPost := new(duoerlapi.Post)
	if post != nil {
		apiPost = &duoerlapi.Post{
			Id:        post.Id.Hex(),
			Content:   template.HTML(post.Content),
			Author:    toApiUser(author),
			CreatedAt: post.CreatedAt.Format(global.CREATED_AT_LONG),
		}
	}
	return apiPost
}
