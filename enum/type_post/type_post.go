package typepost

var (
	Text     = "text"
	Checkbox = "checkbox"
	Image    = "image"
)

func GetAllTypes() []string {
	return []string{Text, Checkbox, Image}
}
