package contenthelper

import (
	"github.com/betarobin/poster/enum/errlist"
	typepost "github.com/betarobin/poster/enum/type_post"
	"github.com/betarobin/poster/model/content"
)

func ParseContent(postType string, postContent interface{}) (interface{}, error) {
	switch postType {
	case typepost.Text:
		return content.ParseText(postContent)
	case typepost.Checklist:
		return content.ParseChecklist(postContent)
	case typepost.Image:
		return content.ParseImage(postContent)
	default:
		return nil, errlist.ErrInvalidPostType
	}
}
