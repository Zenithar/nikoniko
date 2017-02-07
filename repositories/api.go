package repositories

import (
	"zenithar.org/go/nikoniko/domain"

	"zenithar.org/go/common/dao/api"
)

// -----------------------------------------------------------------------------

// UserSearchFilter is the filter used to search in user repository
type UserSearchFilter struct {
	Query string
	Email string
}

// UserRepository is the user repository contract
type UserRepository interface {
	Register(domain.User) (*domain.User, error)
	Get(string) (*domain.User, error)
	Update(string, domain.User) (*domain.User, error)
	Delete(string) (*domain.User, error)
	Search(*UserSearchFilter, *api.SortParameters, *api.Pagination) ([]*domain.User, uint64, error)

	GetByEmail(string) (*domain.User, error)
}

// -----------------------------------------------------------------------------

// VoteSearchFilter is the filter used to search in vote repository
type VoteSearchFilter struct {
	UserID  string
	MinDate int64
	MaxDate int64
}

// VoteRepository is the vote repository contract
type VoteRepository interface {
	Register(domain.Vote) (*domain.Vote, error)
	Get(string) (*domain.Vote, error)
	Update(string, domain.Vote) (*domain.Vote, error)
	Delete(string) (*domain.Vote, error)
	Search(*VoteSearchFilter, *api.SortParameters, *api.Pagination) ([]*domain.Vote, uint64, error)
}
