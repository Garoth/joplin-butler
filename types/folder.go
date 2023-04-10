package types

import "fmt"

type Folder struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	CreatedTime       int    `json:"created_time"`
	UpdatedTime       int    `json:"updated_time"`
	UserCreatedTime   int    `json:"user_created_time"`
	UserUpdatedTime   int    `json:"user_updated_time"`
	EncryptionCipher  string `json:"encryption_cipher_text"`
	EncryptionApplied int    `json:"encryption_applied"`
	ParentID          string `json:"parent_id"`
	IsShared          int    `json:"is_shared"`
	ShareID           string `json:"share_id"`
	MasterKeyID       string `json:"master_key_id"`
	Icon              string `json:"icon"`
}

func (me Folder) String() string {
	return me.ID + " " + me.Title
}

func (me Folder) DetailedString() string {
	out := fmt.Sprintf("ID: %v\n"+
		"Title: %v\n"+
		"ParentID: %v\n"+
		"CreatedTime: %v\n"+
		"UpdatedTime: %v\n"+
		"UserCreatedTime: %v\n"+
		"UserUpdatedTime: %v\n"+
		"IsShared: %v\n"+
		"ShareID: %v\n"+
		"Icon: %v\n",
		me.ID, me.Title, me.ParentID, me.CreatedTime, me.UpdatedTime,
		me.UserCreatedTime, me.UserUpdatedTime,
		me.IsShared, me.ShareID, me.Icon)
	return out
}
