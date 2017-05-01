package models

type DeletefolderResponse struct {
	ID       string `json:"id"`
	Metadata struct {
		Path           string `json:"path"`
		Name           string `json:"name"`
		Modified       string `json:"modified"`
		IsMine         bool   `json:"ismine"`
		Id             string `json:"id"`
		IsShared       bool   `json:"isshared"`
		IsFolder       bool   `json:"isfolder"`
		IsDeleted      bool   `json:"isdeleted"`
		ParentFolderId int    `json:"parentfolderid"`
		FolderId       int    `json:"folderid"`
	} `json:"metadata"`
}

type DeletefolderRecursiveResponse struct {
	DeletedFiles   int `json:"deletedfiles"`
	DeletedFolders int `json:"deletedfolders"`
}
