package content

import (
	"github.com/betarobin/poster/enum/errlist"
)

type Text struct {
	Text string `json:"text"`
}

func ParseText(jsonInterface interface{}) (*Text, error) {
	m, ok := jsonInterface.(map[string]interface{})
	if !ok {
		return nil, errlist.ErrInvalidContent
	}

	text := &Text{}

	if parsedText, ok := m["text"].(string); ok {
		text.Text = parsedText
		return text, nil
	}

	return nil, errlist.ErrInvalidContent
}
