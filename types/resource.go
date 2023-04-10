package types

type Resource struct {
	ID                      string `json:"id"`
	Title                   string `json:"title"`
	MIME                    string `json:"mime"`
	FileName                string `json:"filename"`
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
