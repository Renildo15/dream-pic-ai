package types

import (
	"time"

	"github.com/google/uuid"
)

type ImageStatus int

const (
	ImageStatusPending ImageStatus = iota
	ImageStatusFailed
	ImageStatusCompleted
)

type Image struct {
	ID            int `bun:"id,pk,autoincrement"`
	UserID        uuid.UUID
	Status        ImageStatus
	BatchID       uuid.UUID
	Prompt        string
	ImageLocation string
	Deleted       bool      `bun:",default:false"`
	CreatedAt     time.Time `bun:",default:now()"`
	DeletedAt     time.Time
}
