package domain

import (
	"gopkg.in/go-playground/validator.v9"
	"zenithar.org/go/common/helpers/unique/uniuri"
)

// User is the user information holder
type User struct {
	ID        string `json:"id" bson:"_id" gorethink:"id" validate:"required"`
	GivenName string `json:"gn" bson:"gn" gorethink:"gn" validate:"required"`
	SurName   string `json:"sn" bson:"sn" gorethink:"sn" validate:"required"`
	Email     string `json:"email" bson:"email" gorethink:"email" validate:"required,email"`
}

// NewUser returns a fresh user instance
func NewUser(gn, sn, email string) (*User, error) {
	entity := &User{
		ID:        uniuri.New().Generate(),
		GivenName: gn,
		SurName:   sn,
		Email:     email,
	}

	// Validates model instance
	if _, err := entity.IsValid(); err != nil {
		return nil, err
	}

	return entity, nil
}

// -----------------------------------------------------------------------------

// IsValid returns the validation status
func (u *User) IsValid() (bool, error) {
	err := validator.New().Struct(u)
	return err == nil, err
}
