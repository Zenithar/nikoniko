package user

import (
	"net/http"
	"strings"

	valid "gopkg.in/asaskevich/govalidator.v5"
	"zenithar.org/go/common/dao/api"
	"zenithar.org/go/common/eventbus"

	"zenithar.org/go/nikoniko/domain"
	"zenithar.org/go/nikoniko/dto"
	"zenithar.org/go/nikoniko/dto/mapper"
	"zenithar.org/go/nikoniko/repositories"
	"zenithar.org/go/nikoniko/services"
)

type userService struct {
	users repositories.UserRepository
	bus   eventbus.EventBus
}

// NewService returns a user service instance
func NewService(users repositories.UserRepository, bus eventbus.EventBus) (services.UserService, error) {
	return &userService{
		users: users,
		bus:   bus,
	}, nil
}

// -----------------------------------------------------------------------------

func (s *userService) Register(r dto.UserCreateReq) *dto.SingleUserRes {
	res := &dto.SingleUserRes{}

	// Check mandatory field
	if len(strings.TrimSpace(r.Email)) == 0 {
		res.Error = &dto.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "User email is mandatory !",
		}
		return res
	}

	if !valid.IsEmail(r.Email) {
		res.Error = &dto.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "A valid email is mandatory !",
		}
		return res
	}

	if len(strings.TrimSpace(r.FirstName)) == 0 {
		res.Error = &dto.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "FirstName is mandatory !",
		}
		return res
	}

	if len(strings.TrimSpace(r.LastName)) == 0 {
		res.Error = &dto.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "LastName is mandatory !",
		}
		return res
	}

	// Check User existence
	entity, err := s.users.GetByEmail(r.Email)
	if err != nil && err != api.ErrNoResult {
		res.Error = &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return res
	}
	if entity != nil {
		res.Error = &dto.Error{
			Code:    http.StatusConflict,
			Message: "User already exists !",
		}
		return res
	}

	// Create the User from the request
	model, err := domain.NewUser(r.Email, r.FirstName, r.LastName)
	if err != nil {
		res.Error = &dto.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return res
	}

	saved, err := s.users.Register(*model)
	if err != nil {
		res.Error = &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return res
	}

	// Create the response
	res.Entity = mapper.FromUser(saved)

	// Broadcast event
	s.bus.Publish(services.UserCreated, res)

	return res
}

// Get a user from given identifier
func (s *userService) Get(id string) *dto.SingleUserRes {
	res := &dto.SingleUserRes{}

	// Check mandatory field
	if len(strings.TrimSpace(id)) == 0 {
		res.Error = &dto.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "User identifier is mandatory !",
		}
		return res
	}

	user, err := s.users.Get(id)
	if err != nil {
		res.Error = &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		if err == api.ErrNoResult {
			res.Error.Code = http.StatusNotFound
		}
		return res
	}

	// Create the response
	res.Entity = mapper.FromUser(user)

	return res
}

func (s *userService) Update(r dto.UserCreateReq) *dto.SingleUserRes {
	res := &dto.SingleUserRes{}

	// Check mandatory field
	if len(strings.TrimSpace(r.Id)) == 0 {
		res.Error = &dto.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "User id is mandatory !",
		}
		return res
	}

	// Check User existence
	entity, err := s.users.Get(r.Id)
	if err != nil {
		res.Error = &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		if err == api.ErrNoResult {
			res.Error.Code = http.StatusNotFound
		}
		return res
	}

	// Update User model attributes
	updated := false

	if entity.Email != r.Email && len(strings.TrimSpace(r.Email)) > 0 {
		// Check if email is a valid email
		if !valid.IsEmail(r.Email) {
			res.Error = &dto.Error{
				Code:    http.StatusPreconditionFailed,
				Message: "Email is invalid !",
			}
			return res
		}

		// Check if email is not alredy used in the directory
		other, err := s.users.GetByEmail(r.Email)
		if err != nil && err != api.ErrNoResult {
			res.Error = &dto.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			return res
		}

		// If email is not owned by user
		if other != nil && other.ID != entity.ID {
			res.Error = &dto.Error{
				Code:    http.StatusNotModified,
				Message: "Email already used !",
			}
			return res
		}

		entity.Email = r.Email
		updated = true
		entity.Touch()
	}

	if entity.FirstName != r.FirstName && len(strings.TrimSpace(r.FirstName)) > 0 {
		entity.FirstName = r.FirstName
		updated = true
		entity.Touch()
	}

	if entity.LastName != r.LastName && len(strings.TrimSpace(r.LastName)) > 0 {
		entity.LastName = r.LastName
		updated = true
		entity.Touch()
	}

	// Save it to the database
	if updated {
		// Validate model
		valid, err := entity.IsValid()
		if !valid || err != nil {
			res.Error = &dto.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
			return res
		}

		saved, err := s.users.Update(r.Id, *entity)
		if err != nil {
			res.Error = &dto.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			return res
		}

		res.Entity = mapper.FromUser(saved)
	} else {
		res.Entity = mapper.FromUser(entity)
	}

	// Broadcast event
	s.bus.Publish(services.UserUpdated, res)

	return res
}

// Delete the given user
func (s *userService) Delete(id string) *dto.SingleUserRes {
	res := &dto.SingleUserRes{}

	// Check mandatory field
	if len(strings.TrimSpace(id)) == 0 {
		res.Error = &dto.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "User identifier is mandatory !",
		}
		return res
	}

	// Check account existence
	entity, err := s.users.Get(id)
	if err != nil {
		res.Error = &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		if err == api.ErrNoResult {
			res.Error.Code = http.StatusNotFound
		}
		return res
	}

	if entity != nil {
		_, err := s.users.Delete(entity.ID)
		if err != nil {
			res.Error = &dto.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			return res
		}

		res.Entity = mapper.FromUser(entity)
	}

	// Broadcast event
	s.bus.Publish(services.UserDeleted, res)

	return res
}

// List all users from the database without pagination
func (s *userService) List(r dto.UserSearchReq) *dto.MultipleUserRes {
	res := &dto.MultipleUserRes{}

	sortParams := api.SortConverter(r.Sorts)

	filter := &repositories.UserSearchFilter{
		Query: r.Query,
		Email: r.Email,
	}
	entities, count, err := s.users.Search(filter, sortParams, nil)
	if err != nil && err != api.ErrNoResult {
		res.Error = &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if err == api.ErrNoResult {
		res.Total = 0
		res.Members = make([]*dto.Domain_User, 0)
	} else {
		res.Total = int32(count)
		res.Members = mapper.FromUsers(entities)
	}

	return res
}

// Search for users, returns a paginated collection
func (s *userService) Search(r dto.UserSearchReq) *dto.PaginatedUserRes {
	res := &dto.PaginatedUserRes{}

	sortParams := api.SortConverter(r.Sorts)
	pagination := api.NewPaginator(uint(r.Page), uint(r.PerPage))

	filter := &repositories.UserSearchFilter{
		Query: r.Query,
		Email: r.Email,
	}
	entities, total, err := s.users.Search(filter, sortParams, pagination)
	if err != nil && err != api.ErrNoResult {
		res.Error = &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	pagination.SetTotal(uint(total))
	res.Total = uint32(pagination.Count())

	if r.Page > 1 && r.Page*r.PerPage > uint32(total) {
		res.Error = &dto.Error{
			Code:    http.StatusBadRequest,
			Message: "Page out of bounds !",
		}
		return res
	}

	if err == api.ErrNoResult {
		res.PageCount = 0
		res.PerPage = r.PerPage
		res.CurrentPage = 1
		res.Members = make([]*dto.Domain_User, 0)
	} else {
		res.PageCount = uint32(pagination.NumPages())
		res.PerPage = uint32(pagination.PerPage)
		res.CurrentPage = r.Page
		res.Members = mapper.FromUsers(entities)
	}

	return res
}
