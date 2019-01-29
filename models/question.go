package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
<<<<<<< HEAD
)

type Question struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
=======
	"github.com/gobuffalo/validate/validators"
)

// Question - Represents a question
type Question struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	QuestionText string    `json:"question" db:"question"`
	Editorial    string    `json:"editorial" db:"editorial"`
	HostID       uuid.UUID `json:"host_id" db:"host_id" has_one:"hosts" fk_id:"id"`
	ContestID    uuid.UUID `json:"contest_id" db:"contest_id" has_one:"contests" fk_id:"id"`
>>>>>>> 09bef744bd171c33fbdfb3a579f56ff782b8e7c9
}

// String is not required by pop and may be deleted
func (q Question) String() string {
	jq, _ := json.Marshal(q)
	return string(jq)
}

// Questions is not required by pop and may be deleted
type Questions []Question

// String is not required by pop and may be deleted
func (q Questions) String() string {
	jq, _ := json.Marshal(q)
	return string(jq)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (q *Question) Validate(tx *pop.Connection) (*validate.Errors, error) {
<<<<<<< HEAD
	return validate.NewErrors(), nil
=======
	return validate.Validate(
		&validators.StringIsPresent{Field: q.QuestionText, Name: "QuestionText"},
		&validators.StringIsPresent{Field: q.Editorial, Name: "Editorial"},
	), nil
>>>>>>> 09bef744bd171c33fbdfb3a579f56ff782b8e7c9
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (q *Question) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (q *Question) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
