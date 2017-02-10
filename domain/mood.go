package domain

// Mood is the mood enumeration
type Mood uint64

const (
	// UndefinedMood indicates that the user has not voted
	UndefinedMood Mood = iota // 0
	// AngryMood indicates tha user is angry during the voting day
	AngryMood // 1
	// BadMood indicates that user have bad feelings during the voting day
	BadMood // 2
	// ModerateMood indicates that user has moderate feelings during voting day
	ModerateMood // 3
	// GoodMood indicates that user has good feelings during voting day
	GoodMood // 4
	// HappyMood idicates that the user is happy during the voting day
	HappyMood // 5
)

var moods = []string{"undefined", "angry", "bad", "moderate", "good", "happy"}

// String returns string representation of the given enumeration
func (m Mood) String() string {
	return moods[m]
}
