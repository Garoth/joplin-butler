package types

import (
	"encoding/json"
)

type Paginateable interface {
	Note | Notebook
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
