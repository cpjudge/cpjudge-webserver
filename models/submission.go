package models

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/pkg/errors"
)

// Submission - To store question submissions
type Submission struct {
	ID             uuid.UUID `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	QuestionID     uuid.UUID `json:"question_id" db:"question_id" has_one:"questions" fk_id:"id"`
	UserID         uuid.UUID `json:"user_id" db:"user_id" has_one:"users" fk_id:"id"`
	Language       string    `json:"language" db:"language"`
	Status         int       `json:"status" db:"status"`
	SubmissionFile string    `json:"-" db:"-"`
}

// AfterCreate - to save the file into a directory
func (s *Submission) AfterCreate(tx *pop.Connection) error {
	if s.SubmissionFile == "" {
		return nil
	}
	dir := filepath.Join(".", "uploads", "submissions", s.QuestionID.String())
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create(filepath.Join(dir, s.ID.String()))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.WriteString(f, s.SubmissionFile)
	return err
}

// String is not required by pop and may be deleted
func (s Submission) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Submissions is not required by pop and may be deleted
type Submissions []Submission

// String is not required by pop and may be deleted
func (s Submissions) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Submission) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Submission) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Submission) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
