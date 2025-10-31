package model

import (
	"time"

	"github.com/google/uuid"
)

type Chunk struct {
	Filename   string
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ChunkIndex int
	ChunkData  []byte `gorm:"type:bytea"`
	CreatedAt  time.Time
}

func (Chunk) TableName() string {
	return "chunk"
}

func (c *Chunk) GetId() uuid.UUID {
	return c.ID
}

func CreateChunk(id string, audio []byte, index int) *Chunk {
	return &Chunk{
		Filename:   id,
		ID:         uuid.New(),
		ChunkIndex: index,
		ChunkData:  audio,
		CreatedAt:  time.Now(),
	}
}
