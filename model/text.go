package model

import (
	"time"

	"github.com/google/uuid"
)

type Text struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	VoiceIndex int
	TextIndex  int
	Start      float64
	End        float64
	Content    string
	CreatedAt  time.Time
}

func (Text) TableName() string {
	return "Text"
}

func (f *Text) GetId() uuid.UUID {
	return f.ID
}

func CreateText(id string, voiceIdx, textIdx int, start, end float64, content string) *Text {
	return &Text{
		ID:         uuid.New(),
		VoiceIndex: voiceIdx,
		TextIndex:  textIdx,
		Start:      start,
		End:        end,
		Content:    content,
	}
}
