package enum

import "errors"

//relation status bewteen 2 users
type RelationState int

const (
	E_Default RelationState = iota
	E_Liked
	E_Disliked
	E_Matched
)

var relationStateDesc = []string{
	"default",
	"liked",
	"disliked",
	"matched",
}

func (s *RelationState) String() string {
	if *s >= E_Default && *s <= E_Matched {
		return relationStateDesc[*s]
	}
	return "unknown"
}


func (s *RelationState) MarshalJSON() ([]byte, error) {
	sVal := s.String()
	if sVal == "unknown" {
		return nil, errors.New("unknown")
	}
	return []byte("\""+sVal+"\""), nil
}