package models

// ListfolderResponse contains server response after listfolder call
type ListfolderResponse struct {
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
		Contents       []struct {
			Path           string `json:"path"`
			Name           string `json:"name"`
			Modified       string `json:"modified"`
			IsMine         bool   `json:"ismine"`
			ID             string `json:"id"`
			IsShared       bool   `json:"isshared"`
			IsFolder       bool   `json:"isfolder"`
			ParentFolderID int    `json:"parentfolderid"`
			FolderID       int    `json:"folderid"`
		} `json:"contents"`
	} `json:"metadata"`
}
