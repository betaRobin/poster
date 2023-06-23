package content

import (
	"github.com/betarobin/poster/enum/errlist"
)

type Checklist struct {
	Checklist []*Checkbox `json:"checklist"`
}

type Checkbox struct {
	SortOrder uint8  `json:"sort_order"`
	IsChecked bool   `json:"is_checked"`
	Text      string `json:"text"`
}

func ParseChecklist(jsonInterface interface{}) (*Checklist, error) {
	m, ok := jsonInterface.([]interface{})

	if !ok {
		return nil, errlist.ErrInvalidContent
	}

	checklist := []*Checkbox{}

	for _, data := range m {
		parsedData, ok := data.(map[string]interface{})

		if !ok {
			continue
		}

		checkbox := &Checkbox{}

		// comes in as float64 in the byte[] interface
		if parsedSortOrder, ok := parsedData["sort_order"].(float64); ok {
			checkbox.SortOrder = uint8(parsedSortOrder)
		} else {
			return nil, errlist.ErrInvalidContent
		}

		if parsedIsChecked, ok := parsedData["is_checked"].(bool); ok {
			checkbox.IsChecked = parsedIsChecked
		} else {
			return nil, errlist.ErrInvalidContent
		}

		if parsedText, ok := parsedData["text"].(string); ok {
			checkbox.Text = parsedText
		} else {
			return nil, errlist.ErrInvalidContent
		}

		checklist = append(checklist, checkbox)
	}

	return &Checklist{Checklist: checklist}, nil
}
