package domain

import (
	"errors"
	"time"

	validator "gopkg.in/asaskevich/govalidator.v5"
	"zenithar.org/go/common/helpers/unique/uniuri"
)

// Vote is the user vote holder
type Vote struct {
	ID        string    `json:"id" bson:"_id" gorethink:"id" valid:"required,ascii"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp" gorethink:"timestamp" valid:"numeric,required"`
	Mood      Mood      `json:"mood" bson:"mood" gorethink:"mood" validate:"numeric,required"`
	Comment   string    `json:"comment" bson:"comment" gorethink:"comment" valid:"utf8"`
	UserID    string    `json:"user_id" bson:"user_id" gorethink:"user_id" valid:"required"`
}

// NewVote returns a fresh new vote instance
func NewVote(userID string, mood Mood) (*Vote, error) {
	entity := &Vote{
		ID:        uniuri.New().Generate(),
		Timestamp: time.Now().UTC(),
		Mood:      mood,
		UserID:    userID,
	}

	// Check entity
	// Validates model instance
	if _, err := entity.IsValid(); err != nil {
		return nil, err
	}

	return entity, nil
}

// -----------------------------------------------------------------------------
var (
	// ErrVoteCommentIsMandatory is raised when comment is missing and mood is at Bad level
	ErrVoteCommentIsMandatory = errors.New("comment is mandatory for bad mood")
)

// -----------------------------------------------------------------------------

// IsValid returns the validation status
func (v *Vote) IsValid() (bool, error) {
	if v.Mood == BadMood && validator.IsNull(v.Comment) {
		return false, ErrVoteCommentIsMandatory
	}
	return validator.ValidateStruct(v)
}
