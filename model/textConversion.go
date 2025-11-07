package model

import (
	"time"

	"github.com/google/uuid"
)

type TextConversion struct {
	Filename  string
	Count     int
	CreatedAt time.Time
}

func (TextConversion) TableName() string {
	return "text_conversion"
}

func (f *TextConversion) GetId() uuid.UUID {
	return uuid.Nil
}

func CreateTextConversion(filename string, count int) *TextConversion {
	return &TextConversion{
		Filename: filename,
		Count:    count,
	}
}
