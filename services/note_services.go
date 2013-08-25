package services

import (
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/models/articles"
	"github.com/kobeld/duoerl/models/notes"
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
	"html/template"
	"labix.org/v2/mgo/bson"
)

func NewNote() (noteInput *duoerlapi.NoteInput) {
	noteInput = &duoerlapi.NoteInput{
		Id: bson.NewObjectId().Hex(),
	}
	return
}

func GetUserNotes(userIdHex string) (apiNotes []*duoerlapi.Note, err error) {

	userId, err := utils.ToObjectId(userIdHex)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	notes, err := notes.FindSomeByUserId(userId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	for _, note := range notes {
		apiNotes = append(apiNotes, toApiNote(note, nil))
	}

	return
}

func CreateNote(input *duoerlapi.NoteInput) (originInput *duoerlapi.NoteInput, err error) {
	originInput = input

	noteId, err := utils.ToObjectId(input.Id)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	authorId, err := utils.ToObjectId(input.AuthorId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	note := notes.Note{
		Id:      noteId,
		Article: *articles.NewArticle(input.Title, input.Content, authorId),
	}

	if err = note.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func ShowNote(noteIdHex, userIdHex string) (apiNote *duoerlapi.Note, err error) {
	noteId, err := utils.ToObjectId(noteIdHex)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	note, err := notes.FindById(noteId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	author, err := users.FindById(note.AuthorId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiNote = toApiNote(note, author)

	return
}

// ----- Private -----

func toApiNote(note *notes.Note, author *users.User) *duoerlapi.Note {
	apiNote := new(duoerlapi.Note)
	if note != nil {
		apiNote = &duoerlapi.Note{
			Id:        note.Id.Hex(),
			Title:     note.Title,
			Content:   template.HTML(note.Content),
			Author:    toApiUser(author),
			Link:      note.Link(),
			CreatedAt: note.CreatedAt.Format(global.CREATED_AT_LONG),
		}
	}
	return apiNote
}
