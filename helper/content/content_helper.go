package contenthelper

import (
	"github.com/betarobin/poster/enum/errlist"
	typepost "github.com/betarobin/poster/enum/type_post"
	"github.com/betarobin/poster/model/content"
)

func ParseContent(postType string, postContent string) (interface{}, error) {
	if len(postContent) == 0 {
		return nil, errlist.ErrInvalidContent
	}

	switch postType {
	case typepost.Text:
		return content.ParseText(postContent)
	case typepost.Checkbox:
		return content.ParseChecklist(postContent)
	case typepost.Image:
		return content.ParseImage(postContent)
	default:
		return nil, errlist.ErrInvalidPostType
	}
}
