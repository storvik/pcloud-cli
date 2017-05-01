package models

type RenamefolderResponse struct {
	Metadata struct {
		Path           string `json:"path"`
		Name           string `json:"name"`
		Modified       string `json:"modified"`
		IsMine         bool   `json:"ismine"`
		Id             string `json:"id"`
		IsShared       bool   `json:"isshared"`
		IsFolder       bool   `json:"isfolder"`
		ParentFolderId int    `json:"parentfolderid"`
		FolderId       int    `json:"folderid"`
	} `json:"metadata"`
}
