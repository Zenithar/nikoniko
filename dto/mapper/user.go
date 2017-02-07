package mapper

import (
	"zenithar.org/go/nikoniko/domain"
	"zenithar.org/go/nikoniko/dto"
)

// FromUser returns a DTO object from domain one
func FromUser(entity *domain.User) *dto.Domain_User {
	return &dto.Domain_User{
		Id:        entity.ID,
		Email:     entity.Email,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
	}
}

// FromUsers returns a DTO Collection from domain one
func FromUsers(entities []*domain.User) []*dto.Domain_User {
	res := []*dto.Domain_User{}
	for _, entity := range entities {
		res = append(res, FromUser(entity))
	}
	return res
}
