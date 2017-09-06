package models

// UploadfileResponse contains server response after update call
type UploadfileResponse struct {
	FiledIds []int `json:"fileids"`
	Metadata []struct {
		Name           string `json:"name"`
		Path           string `json:"path"`
		Category       int    `json:"category"`
		Modified       string `json:"modified"`
		ID             string `json:"id"`
		FileID         int    `json:"fileid"`
		IsFolder       bool   `json:"isfolder"`
		ParentFolderID int    `json:"parentfolderid"`
		Size           int    `json:"size"`
		ContentType    string `json:"contenttype"`
	} `json:"metadata"`
}
