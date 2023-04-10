package types

import (
	"fmt"
	"encoding/json"
)

type Paginateable interface {
	ItemInfo | Note
}

type Paginated[P Paginateable] struct {
	Items []P `json:"items"`
	HasMore bool `json:"has_more"`
}

func NewPaginated[P Paginateable](jsonStr string) (*Paginated[P], error) {
	me := &Paginated[P]{}
	me.Items = []P{}

	if len(jsonStr) > 0 {
		err := json.Unmarshal([]byte(jsonStr), me)
		if err != nil  {
			return nil, err
		}
	}

	return me, nil
}

func (me *Paginated[P]) String() string {
	output := ""
	for i, item := range me.Items {
		if i == len(me.Items) - 1 {
			output += fmt.Sprintf("%v", item)
		} else {
			output += fmt.Sprintf("%v\n", item)
		}
	}
	return output
}
