package note

import (
	"chat-service/internal/converter"
	"chat-service/pkg/note_v1"
	"context"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *note_v1.CreateRequest) (*note_v1.CreateResponse, error) {
	id, err := i.noteService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted note with id: %d", id)

	return &note_v1.CreateResponse{
		Id: id,
	}, nil
}
