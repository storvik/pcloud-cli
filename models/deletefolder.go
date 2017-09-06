package models

// DeletefolderResponse contains server response after deletefolder call
type DeletefolderResponse struct {
	ID       string `json:"id"`
	Metadata struct {
		Path           string `json:"path"`
		Name           string `json:"name"`
		Modified       string `json:"modified"`
		IsMine         bool   `json:"ismine"`
		ID             string `json:"id"`
		IsShared       bool   `json:"isshared"`
		IsFolder       bool   `json:"isfolder"`
		IsDeleted      bool   `json:"isdeleted"`
		ParentFolderID int    `json:"parentfolderid"`
		FolderID       int    `json:"folderid"`
	} `json:"metadata"`
}

// DeletefolderRecursiveResponse contains server response after recursive delete folder call
type DeletefolderRecursiveResponse struct {
	DeletedFiles   int `json:"deletedfiles"`
	DeletedFolders int `json:"deletedfolders"`
}
