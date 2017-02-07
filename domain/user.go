package domain

import (
	"time"

	validator "gopkg.in/asaskevich/govalidator.v5"
	"zenithar.org/go/common/helpers/unique/uniuri"
)

// User is the user information holder
type User struct {
	ID             string    `json:"id" bson:"_id" gorethink:"id" valid:"required"`
	FirstName      string    `json:"first_name" bson:"first_name" gorethink:"first_name" valid:"required"`
	LastName       string    `json:"last_name" bson:"last_name" gorethink:"last_name" valid:"required"`
	Email          string    `json:"email" bson:"email" gorethink:"email" valid:"required,email"`
	LastModifiedAt time.Time `json:"last_modified_at" bson:"last_modified_at" gorethink:"last_modified_at"`
}

// NewUser returns a fresh user instance
func NewUser(gn, sn, email string) (*User, error) {
	entity := &User{
		ID:             uniuri.New().Generate(),
		FirstName:      gn,
		LastName:       sn,
		Email:          email,
		LastModifiedAt: time.Now().UTC(),
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
	return validator.ValidateStruct(u)
}

// Touch updates the last modified date
func (u *User) Touch() {
	u.LastModifiedAt = time.Now().UTC()
}
