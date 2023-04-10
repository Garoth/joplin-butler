package types

type Event struct {
	ID               int    `json:"id"`
	ItemType         int    `json:"item_type"`
	ItemID           string `json:"item_id"`
	Type             int    `json:"type"`
	CreatedTime      int    `json:"created_time"`
	Source           int    `json:"source"`
	BeforeChangeItem string `json:"before_change_item"`
}
