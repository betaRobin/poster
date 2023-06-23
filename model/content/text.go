package content

import (
	"encoding/json"

	"github.com/betarobin/poster/enum/errlist"
)

type Text struct {
	Text string `json:"text"`
}

func ParseText(jsonString string) (*Text, error) {
	text := &Text{}

	err := json.Unmarshal([]byte(jsonString), text)
	if err != nil {
		return nil, errlist.ErrInvalidContent
	}

	return text, nil
}
