package handler

import (
	"github.com/micro/go-micro/v2/client"

	notes "github.com/micro/services/notes/proto"
)

// NewHandler returns an initialized Handler
func NewHandler(client client.Client) *Handler {
	return &Handler{
		notes: notes.NewNotesService("go.micro.srv.notes", client),
	}
}

// Handler imlements the notes proto definition
type Handler struct {
	notes notes.NotesService
}
