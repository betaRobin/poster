package contenthelper

import (
	"encoding/json"

	"github.com/betarobin/poster/enum/errlist"
	typepost "github.com/betarobin/poster/enum/type_post"
	"github.com/betarobin/poster/model/content"
)

func ParseContent(postType string, postContent interface{}) (interface{}, error) {
	switch postType {
	case typepost.Text:
		return ParseText(postContent)
	case typepost.Checklist:
		return ParseChecklist(postContent)
	case typepost.Image:
		return ParseImage(postContent)
	default:
		return nil, errlist.ErrInvalidPostType
	}
}

func ParseChecklist(jsonInterface interface{}) (*content.Checklist, error) {
	m, ok := jsonInterface.(map[string]interface{})

	if !ok {
		return nil, errlist.ErrInvalidContent
	}

	checklist, ok := m["checklist"].([]interface{})

	if !ok {
		return nil, errlist.ErrInvalidContent
	}

	parsedChecklist := []*content.Checkbox{}

	for _, data := range checklist {
		parsedData, ok := data.(map[string]interface{})

		if !ok {
			continue
		}

		checkbox := &content.Checkbox{}

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

		parsedChecklist = append(parsedChecklist, checkbox)
	}

	return &content.Checklist{Checklist: parsedChecklist}, nil
}

func ParseImage(jsonInterface interface{}) (*content.Image, error) {
	m, ok := jsonInterface.(map[string]interface{})

	if !ok {
		return nil, errlist.ErrInvalidContent
	}

	image := &content.Image{}

	if parsedImage, ok := m["image"].(string); ok {
		image.Image = parsedImage
	} else {
		return nil, errlist.ErrInvalidContent
	}

	if parsedText, ok := m["text"].(string); ok {
		image.Text = parsedText
	} else {
		return nil, errlist.ErrInvalidContent
	}

	return image, nil
}

func ParseText(jsonInterface interface{}) (*content.Text, error) {
	m, ok := jsonInterface.(map[string]interface{})
	if !ok {
		return nil, errlist.ErrInvalidContent
	}

	text := &content.Text{}

	if parsedText, ok := m["text"].(string); ok {
		text.Text = parsedText
		return text, nil
	}

	return nil, errlist.ErrInvalidContent
}

func ContentToJsonString(postType string, contentInterface interface{}) (string, error) {
	parsedContent, _ := ParseContent(postType, contentInterface)
	contentJsonString := ``
	switch postType {
	case typepost.Text:
		typedContent := parsedContent.(*content.Text)
		contentJson, _ := json.Marshal(typedContent)
		contentJsonString = string(contentJson)
	case typepost.Image:
		typedContent := parsedContent.(*content.Image)
		contentJson, _ := json.Marshal(typedContent)
		contentJsonString = string(contentJson)
	case typepost.Checklist:
		typedContent := parsedContent.(*content.Checklist)
		contentJson, _ := json.Marshal(typedContent)
		contentJsonString = string(contentJson)
	default:
		return "", errlist.ErrInvalidContent
	}

	return contentJsonString, nil
}
