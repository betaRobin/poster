package content

import (
	"github.com/betarobin/poster/enum/errlist"
)

type Image struct {
	Image string `json:"image"` // base64 string
	Text  string `json:"text"`
}

func ParseImage(jsonInterface interface{}) (*Image, error) {
	m, ok := jsonInterface.(map[string]interface{})

	if !ok {
		return nil, errlist.ErrInvalidContent
	}

	image := &Image{}

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

	/*
		m, ok := jsonInterface.(map[string]interface{})
		if !ok {
			return nil, errlist.ErrInvalidContent
		}

		text := &Text{}

		if parsedText, ok := m["text"].(string); ok {
			fmt.Println(parsedText)
			text.Text = parsedText
			return text, nil
		}

		return nil, errlist.ErrInvalidContent
	*/
}
