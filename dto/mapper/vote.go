package mapper

import (
	"zenithar.org/go/nikoniko/domain"
	"zenithar.org/go/nikoniko/dto"
)

// FromVote returns a DTO object from domain one
func FromVote(entity *domain.Vote) *dto.Domain_Vote {
	return &dto.Domain_Vote{
		Id:        entity.ID,
		Timestamp: entity.Timestamp.Unix(),
		Comment:   entity.Comment,
		UserId:    entity.UserID,
		Mood:      FromMood(entity.Mood).String(),
	}
}

// FromVotes returns a DTO Collection from domain one
func FromVotes(entities []*domain.Vote) []*dto.Domain_Vote {
	res := []*dto.Domain_Vote{}
	for _, entity := range entities {
		res = append(res, FromVote(entity))
	}
	return res
}
