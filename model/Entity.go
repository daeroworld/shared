package model

import "github.com/google/uuid"

type Entity interface {
	GetId() uuid.UUID
}
