package routes

import (
	"net/http"

	"zenithar.org/go/nikoniko/services"
)

// VoteController is the vote form controller
type VoteController struct {
	voteService services.VoteService
}

// NewVoteController returns a vote controller instance
func NewVoteController(voteService services.VoteService) (*VoteController, error) {
	return &VoteController{
		voteService: voteService,
	}, nil
}

// -----------------------------------------------------------------------------

// Index returns the index page
func (c *VoteController) Index(w http.ResponseWriter, r *http.Request) {
	// Retreive content object from context
	content, _ := FromContext(r.Context())

	// Prepare the view model
	content.Data = map[string]interface{}{
		"Title": "Welcome",
	}
	content.Template = "welcome"
}
