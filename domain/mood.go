package domain

// Mood is the mood enumeration
type Mood uint64

const (
	// UndefinedMood indicates that the user has not voted
	UndefinedMood Mood = iota // 0
	// BadMood indicates that user have bad feelings during the voting day
	BadMood // 1
	// ModerateMood indicates that user has moderate feelings during voting day
	ModerateMood // 2
	// GoodMood indicates that user has good feelings during voting day
	GoodMood // 3
)

var moods = []string{"undefined", "bad", "moderate", "good"}

// String returns string representation of the given enumeration
func (m Mood) String() string {
	return moods[m]
}
