package mapper

import (
	"zenithar.org/go/nikoniko/domain"
	"zenithar.org/go/nikoniko/dto"
)

// FromMood converts a domain mood to dto one
func FromMood(entity domain.Mood) dto.Domain_Mood {
	switch entity {
	case domain.BadMood:
		return dto.Domain_Bad
	case domain.ModerateMood:
		return dto.Domain_Moderate
	case domain.GoodMood:
		return dto.Domain_Good
	default:
		return dto.Domain_Undefined
	}
}
