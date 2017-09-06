package models

// CreatefolderResponse contains server response after createfolder call
type CreatefolderResponse struct {
	Metadata struct {
		Path           string `json:"path"`
		Name           string `json:"name"`
		Modified       string `json:"modified"`
		IsMine         bool   `json:"ismine"`
		ID             string `json:"id"`
		IsShared       bool   `json:"isshared"`
		IsFolder       bool   `json:"isfolder"`
		ParentFolderID int    `json:"parentfolderid"`
		FolderID       int    `json:"folderid"`
	} `json:"metadata"`
}
