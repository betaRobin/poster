package content

import (
	"encoding/json"

	"github.com/betarobin/poster/enum/errlist"
)

type Checkboxes struct {
	Checkboxes []*Checkbox `json:"checkboxes"`
}

type Checkbox struct {
	SortOrder uint8  `json:"sort_order"`
	IsChecked bool   `json:"is_checked"`
	Text      string `json:"text"`
}

func ParseCheckboxes(jsonString string) (*Checkboxes, error) {
	checkboxes := &Checkboxes{}

	err := json.Unmarshal([]byte(jsonString), checkboxes)
	if err != nil {
		return nil, errlist.ErrInvalidContent
	}

	return checkboxes, nil
}

/*
Marshalling/Unmarshalling example

func main() {
	cbox1 := &Checkbox{SortOrder: 1, IsChecked: false, Text: "false"}
	cbox2 := &Checkbox{SortOrder: 2, IsChecked: true, Text: "true"}

	cboxslice := []*Checkbox{cbox1, cbox2}
	cboxes := &Checkboxes{Data: cboxslice}
	cboxjson, _ := json.Marshal(cboxes)

	fmt.Println("Marshalled: " + string(cboxjson))

	parsedCbox, err := ParseCheckboxes(string(cboxjson))

	if err != nil {
		fmt.Println("Failed to parse jsong string")
	} else {
		for _, item := range parsedCbox.Data {
			fmt.Printf("SortOrder : %d\nIsChecked : %t\nText : %s\n\n", item.SortOrder, item.IsChecked, item.Text)
		}
	}
}
*/
