// The various "item types" that exist in the Joplin database
// This relates to https://joplinapp.org/api/references/rest_api/#item-type-ids
package types

import (
	"fmt"
	"strings"
)

type ItemTypeID int

const (
	ItemTypeNone = iota
	ItemTypeNote
	ItemTypeFolder
	ItemTypeSetting
	ItemTypeResource
	ItemTypeTag
	ItemTypeNoteTag
	ItemTypeSearch
	ItemTypeAlarm
	ItemTypeMasterKey
	ItemTypeItemChange
	ItemTypeNoteResource
	ItemTypeResourceLocalState
	ItemTypeRevision
	ItemTypeMigration
	ItemTypeSmartFilter
	ItemTypeCommand
)

var (
	ItemTypeMap = map[string]ItemTypeID{
		"note":                 ItemTypeNote,
		"folder":               ItemTypeFolder,
		"setting":              ItemTypeSetting,
		"resource":             ItemTypeResource,
		"tag":                  ItemTypeTag,
		"note_tag":             ItemTypeNoteTag,
		"search":               ItemTypeSearch,
		"alarm":                ItemTypeAlarm,
		"master_key":           ItemTypeMasterKey,
		"item_change":          ItemTypeItemChange,
		"note_resource":        ItemTypeNoteResource,
		"resource_local_state": ItemTypeResourceLocalState,
		"revision":             ItemTypeRevision,
		"migration":            ItemTypeMigration,
		"smart_filter":         ItemTypeSmartFilter,
		"command":              ItemTypeCommand,
	}
)

func NewItemTypeID(itemTypeStr string) (ItemTypeID, error) {
	itemType, ok := ItemTypeMap[strings.ToLower(itemTypeStr)]
	if !ok {
		return ItemTypeNone, fmt.Errorf("Unknown item type '%s'", itemTypeStr)
	}
	return itemType, nil
}

func (me ItemTypeID) String() string {
	for k, v := range ItemTypeMap {
		if me == v {
			return k
		}
	}
	return ""
}

type ItemInfo struct {
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
	Title    string `json:"title"`
}

func (me ItemInfo) String() string {
	return me.ID + " " + me.Title
}

func (me ItemInfo) DetailedString() string {
	out := fmt.Sprintf("ID: %v\n"+
		"Title: %v\n"+
		"ParentID: %v"+
		me.ID, me.Title, me.ParentID)
	return out
}
