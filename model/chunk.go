package model

import (
	"time"

	"github.com/google/uuid"
)

type Chunk struct {
	Filename  string
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Index     int
	Data      []byte `gorm:"type:bytea"`
	CreatedAt time.Time
}

func (Chunk) TableName() string {
	return "chunk"
}

func (c *Chunk) GetId() uuid.UUID {
	return c.ID
}

func CreateChunk(id string, audio []byte, index int) *Chunk {
	return &Chunk{
		Filename:  id,
		ID:        uuid.New(),
		Index:     index,
		Data:      audio,
		CreatedAt: time.Now(),
	}
}
