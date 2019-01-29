package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// TestCase - Test case for each question
type TestCase struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	TestCaseLink string    `json:"test_case" db:"test_case"`
	QuestionID   uuid.UUID `json:"question_id" db:"question_id" has_one:"questions" fk_id:"id"`
}

// String is not required by pop and may be deleted
func (t TestCase) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// TestCases is not required by pop and may be deleted
type TestCases []TestCase

// String is not required by pop and may be deleted
func (t TestCases) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *TestCase) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.TestCaseLink, Name: "TestCaseLink"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *TestCase) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *TestCase) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
