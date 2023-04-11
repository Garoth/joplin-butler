package types

import "fmt"

type Resource struct {
	ID                      string `json:"id"`
	Title                   string `json:"title"`
	MIME                    string `json:"mime"`
	Filename                string `json:"filename"`
	CreatedTime             int    `json:"created_time"`
	UpdatedTime             int    `json:"updated_time"`
	UserCreatedTime         int    `json:"user_created_time"`
	UserUpdatedTime         int    `json:"user_updated_time"`
	FileExtension           string `json:"file_extension"`
	EncryptionCipher        string `json:"encryption_cipher_text"`
	EncryptionApplied       int    `json:"encryption_applied"`
	EncryptionBlobEncrypted int    `json:"encryption_blob_encrypted"`
	Size                    int    `json:"size"`
	IsShared                int    `json:"is_shared"`
	ShareID                 string `json:"share_id"`
	MasterKeyID             string `json:"master_key_id"`
}

func (me Resource) String() string {
	return me.ID + " " + me.Filename + " " + me.Title
}

func (me Resource) DetailedString() string {
	out := fmt.Sprintf("ID: %v\n"+
		"Title: %v\n"+
		"Filename: %v\n"+
		"CreatedTime: %v\n"+
		"UpdatedTime: %v\n"+
		"UserCreatedTime: %v\n"+
		"UserUpdatedTime: %v\n"+
		"Size: %v\n"+
		"IsShared: %v\n"+
		"ShareID: %v",
		me.ID, me.Title, me.Filename, me.CreatedTime,
		me.UpdatedTime, me.UserCreatedTime, me.UserUpdatedTime, me.Size,
		me.IsShared, me.ShareID)
	return out
}
