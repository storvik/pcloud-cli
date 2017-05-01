package models

type ChecksumfileResponse struct {
	Sha      string `json:"sha1"`
	Md       string `json:"md5"`
	Metadata struct {
		Category       int    `json:"category"`
		Name           string `json:"name"`
		Modified       string `json:"modified"`
		IsMine         bool   `json:"ismine"`
		Id             string `json:"id"`
		IsShared       bool   `json:"isshared"`
		IsFolder       bool   `json:"isfolder"`
		ParentFolderId int    `json:"parentfolderid"`
		FolderId       int    `json:"folderid"`
		Size           int    `json:"size"`
		ContentType    string `json:"contenttype"`
	} `json:"metadata"`
}
