package content

import (
	"encoding/json"

	"github.com/betarobin/poster/enum/errlist"
)

type Image struct {
	Image string `json:"image"` // base64 string
	Text  string `json:"text"`
}

func ParseImage(jsonString string) (*Image, error) {
	image := &Image{}

	err := json.Unmarshal([]byte(jsonString), image)
	if err != nil {
		return nil, errlist.ErrInvalidContent
	}

	return image, nil
}
