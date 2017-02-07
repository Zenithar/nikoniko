package services

import "zenithar.org/go/nikoniko/dto"

// UserService represents the user service contract
type UserService interface {
	Register(dto.UserCreateReq) *dto.SingleUserRes
	Get(string) *dto.SingleUserRes
	Update(dto.UserCreateReq) *dto.SingleUserRes
	Delete(string) *dto.SingleUserRes
	List(dto.UserSearchReq) *dto.MultipleUserRes
	Search(dto.UserSearchReq) *dto.PaginatedUserRes
}

// VoteService represents the vote service contract
type VoteService interface {
	Register(dto.VoteCreateReq) *dto.SingleVoteRes
	List(dto.VoteSearchReq) *dto.MultipleVoteRes
	Search(dto.VoteSearchReq) *dto.PaginatedVoteRes
}
