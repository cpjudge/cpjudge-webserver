package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

// ParticipateIn - Represents the contests in which the user has taken part and
// vice-versa
type ParticipateIn struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id" has_one:"users" fk_id:"id"`
	ContestID uuid.UUID `json:"contest_id" db:"contest_id" has_one:"contests" fk_id:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p ParticipateIn) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// TableName overrides the table name used by Pop.
func (p ParticipateIn) TableName() string {
	return "participate_in"
}

// ParticipateIns is not required by pop and may be deleted
type ParticipateIns []ParticipateIn

// String is not required by pop and may be deleted
func (p ParticipateIns) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *ParticipateIn) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *ParticipateIn) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *ParticipateIn) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
