package types

type Tag struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	CreatedTime       int    `json:"created_time"`
	UpdatedTime       int    `json:"updated_time"`
	UserCreatedTime   int    `json:"user_created_time"`
	UserUpdatedTime   int    `json:"user_updated_time"`
	EncryptionCipher  string `json:"encryption_cipher_text"`
	EncryptionApplied int    `json:"encryption_applied"`
	IsShared          int    `json:"is_shared"`
	ParentID          string `json:"parent_id"`
}
