package content

type Checklist struct {
	Checklist []*Checkbox `json:"checklist"`
}

type Checkbox struct {
	SortOrder uint8  `json:"sort_order"`
	IsChecked bool   `json:"is_checked"`
	Text      string `json:"text"`
}
