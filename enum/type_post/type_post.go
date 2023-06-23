package typepost

var (
	Text      = "text"
	Checklist = "checklist"
	Image     = "image"
)

func GetAllTypes() []string {
	return []string{Text, Checklist, Image}
}
