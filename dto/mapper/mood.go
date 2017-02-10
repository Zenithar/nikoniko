package mapper

import (
	"zenithar.org/go/nikoniko/domain"
	"zenithar.org/go/nikoniko/dto"
)

var (
	moodMap = map[domain.Mood]dto.Domain_Mood{
		domain.AngryMood:    dto.Domain_Angry,
		domain.BadMood:      dto.Domain_Bad,
		domain.ModerateMood: dto.Domain_Moderate,
		domain.GoodMood:     dto.Domain_Good,
		domain.HappyMood:    dto.Domain_Happy,
	}
)

// FromMood converts a domain mood to dto one
func FromMood(entity domain.Mood) dto.Domain_Mood {
	if value, ok := moodMap[entity]; ok {
		return value
	}
	return dto.Domain_Undefined
}

// ToMood converts a dto mood to domain one
func ToMood(object dto.Domain_Mood) domain.Mood {
	for k, v := range moodMap {
		if v == object {
			return k
		}
	}
	return domain.UndefinedMood
}
