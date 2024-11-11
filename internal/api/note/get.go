package note

import (
	"chat-service/internal/converter"
	"chat-service/pkg/note_v1"
	"context"
	"log"
)

func (i *Implementation) Get(ctx context.Context, req *note_v1.GetRequest) (*note_v1.GetResponse, error) {
	noteObj, err := i.noteService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, title: %s, body: %s, created_at: %v, updated_at: %v\n", noteObj.ID, noteObj.Info.Title, noteObj.Info.Content, noteObj.CreatedAt, noteObj.UpdatedAt)

	return &note_v1.GetResponse{
		Note: converter.ToNoteFromService(noteObj),
	}, nil
}
