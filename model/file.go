package model

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Filename    string
	DurationMs  int64
	Format      string
	SizeBytes   int64
	TotalChunks int
	Status      string                 `gorm:"default:pending"`
	Config      map[string]interface{} `gorm:"type:jsonb"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (File) TableName() string {
	return "file"
}

func (f *File) GetId() uuid.UUID {
	return f.ID
}

func CreateFile(id string, size int, totalchunks int) *File {
	return &File{
		ID:          uuid.New(),
		Filename:    id,
		DurationMs:  12500, // 2minutes 5 seconds
		Format:      "wav",
		SizeBytes:   int64(size),
		TotalChunks: totalchunks,
		Status:      "completed",
		CreatedAt:   time.Now().Add(-5 * time.Minute),
		UpdatedAt:   time.Now(),
	}
}

func NewMockFile() *File {

	return &File{
		ID:          uuid.New(),
		Filename:    "ff",
		DurationMs:  12500, // 2minutes 5 seconds
		Format:      "wav",
		SizeBytes:   int64(66),
		TotalChunks: 2,
		Status:      "completed",
		CreatedAt:   time.Now().Add(-5 * time.Minute),
		UpdatedAt:   time.Now(),
	}
}
