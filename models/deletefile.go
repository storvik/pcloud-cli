package models

type DeletefileResponse struct {
	Metadata struct {
		Category       int    `json:"category"`
		Name           string `json:"name"`
		Modified       string `json:"modified"`
		IsMine         bool   `json:"ismine"`
		Id             string `json:"id"`
		IsShared       bool   `json:"isshared"`
		IsFolder       bool   `json:"isfolder"`
		IsDeleted      bool   `json:"isdeleted"`
		ParentFolderId int    `json:"parentfolderid"`
		Size           int    `json:"size"`
		ContentType    string `json:"contenttype"`
		FileID         int    `json:"fileid"`
	} `json:"metadata"`
}
