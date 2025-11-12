package model

import (
	"time"

	"github.com/google/uuid"
)

type Text struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Filename    string
	ChunkIndex  int
	TextIndex   int
	Index       int
	Start       float64
	End         float64
	Content     string
	EditContent string
	CreatedAt   time.Time
}

func (Text) TableName() string {
	return "text"
}

func (f *Text) GetId() uuid.UUID {
	return f.ID
}

func CreateText(filename string, chunkIdx, textIdx, idx int, start, end float64, content string) *Text {
	return &Text{
		ID:          uuid.New(),
		Filename:    filename,
		ChunkIndex:  chunkIdx,
		TextIndex:   textIdx,
		Index:       idx,
		Start:       start,
		End:         end,
		Content:     content,
		EditContent: "",
	}
}
