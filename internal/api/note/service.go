package note

import (
	"chat-service/internal/service"
	"chat-service/pkg/note_v1"
)

type Implementation struct {
	note_v1.UnimplementedNoteV1Server
	noteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{noteService: noteService}

}
