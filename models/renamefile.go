package models

// RenamefileResponse contains server response after rename file call
type RenamefileResponse struct {
	Metadata struct {
		Category       int    `json:"category"`
		Name           string `json:"name"`
		Modified       string `json:"modified"`
		IsMine         bool   `json:"ismine"`
		ID             string `json:"id"`
		FileID         int    `json:"fileid"`
		IsShared       bool   `json:"isshared"`
		IsFolder       bool   `json:"isfolder"`
		ParentFolderID int    `json:"parentfolderid"`
		FolderID       int    `json:"folderid"`
		Size           int    `json:"size"`
		ContentType    string `json:"contenttype"`
	} `json:"metadata"`
}
