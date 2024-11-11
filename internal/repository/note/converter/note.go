package converter

import (
	modelRepo "chat-service/internal/repository/note/model"
)

func ToNoteFromRepo(note *modelRepo.Note) *modelRepo.Note {
	return &modelRepo.Note{
		ID:        note.ID,
		Info:      ToNoteInfoFromRepo(note.Info),
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}

func ToNoteInfoFromRepo(info modelRepo.NoteInfo) modelRepo.NoteInfo {
	return modelRepo.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
