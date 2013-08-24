package notes

import (
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"github.com/theplant/formdata"
	"net/http"
)

var (
	noteFields = []string{"Id", "Title", "Content"}
)

type NoteViewData struct {
	NoteInput *duoerlapi.NoteInput
	ApiNote   *duoerlapi.Note
	Validated *govalidations.Validated
}

func New(env Env) (status Status, headers Headers, body Body) {

	noteInput := services.NewNote()

	noteViewData := &NoteViewData{
		NoteInput: noteInput,
	}

	mangotemplate.ForRender(env, "notes/new", noteViewData)
	return
}

func Create(env Env) (status Status, headers Headers, body Body) {

	noteInput := new(duoerlapi.NoteInput)
	formdata.UnmarshalByNames(env.Request().Request, &noteInput, noteFields)
	noteInput.AuthorId = services.FetchUserIdFromSession(env)

	result, err := services.CreateNote(noteInput)
	if validated, ok := err.(*govalidations.Validated); ok {
		viewData := &NoteViewData{
			NoteInput: noteInput,
			Validated: validated,
		}
		mangotemplate.ForRender(env, "notes/new", viewData)
		return
	}
	if err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/note/"+result.Id)
}

func Show(env Env) (status Status, headers Headers, body Body) {
	noteId := env.Request().URL.Query().Get(":id")
	currentUserId := services.FetchUserIdFromSession(env)

	apiNote, err := services.ShowNote(noteId, currentUserId)
	if err != nil {
		panic(err)
	}

	noteViewData := &NoteViewData{
		ApiNote: apiNote,
	}

	mangotemplate.ForRender(env, "notes/show", noteViewData)
	return
}
