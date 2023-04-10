package types

import "fmt"

type Note struct {
	ID                   string `json:"id"`
	ParentID             string `json:"parent_id"`
	Title                string `json:"title"`
	Body                 string `json:"body"`
	CreatedTime          int    `json:"created_time"`
	UpdatedTime          int    `json:"updated_time"`
	IsConflict           int    `json:"is_conflict"`
	Latitude             string `json:"latitude"`
	Longitude            string `json:"longitude"`
	Altitude             string `json:"altitude"`
	Author               string `json:"author"`
	SourceURL            string `json:"source_url"`
	IsTodo               int    `json:"is_todo"`
	TodoDue              int    `json:"todo_due"`
	TodoCompleted        int    `json:"todo_completed"`
	Source               string `json:"source"`
	SourceApplication    string `json:"source_application"`
	ApplicationData      string `json:"application_data"`
	Order                int    `json:"order"`
	UserCreatedTime      int    `json:"user_created_time"`
	UserUpdatedTime      int    `json:"user_updated_time"`
	EncryptionCipherText string `json:"encryption_cipher_text"`
	EncryptionApplied    int    `json:"encryption_applied"`
	MarkupLanguage       int    `json:"markup_language"`
	IsShared             int    `json:"is_shared"`
	ShareID              string `json:"share_id"`
	ConflictOriginalID   string `json:"conflict_original_id"`
	MasterKeyID          string `json:"master_key_id"`
	BodyHTML             string `json:"body_html"`
	BaseURL              string `json:"base_url"`
	ImageDataURL         string `json:"image_data_url"`
	CropRect             string `json:"crop_rect"`
}

func (me Note) String() string {
	out := me.ID + " " + me.Title
	if len(out) > 76 {
		out = out[0:76]
	}
	return out
}

func (me Note) DetailedString() string {
	out := fmt.Sprintf("ID: %v\n"+
		"Title: %v\n"+
		"ParentID: %v\n"+
		"CreatedTime: %v\n"+
		"UpdatedTime: %v\n"+
		"Source: %v\n"+
		"---\n%v",
		me.ID, me.Title, me.ParentID, me.CreatedTime, me.UpdatedTime, me.Source, me.Body)
	return out
}
