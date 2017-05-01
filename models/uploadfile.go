package models

type UploadfileResponse struct {
	FiledIds []int `json:"fileids"`
	Metadata []struct {
		Name           string `json:"name"`
		Path           string `json:"path"`
		Category       int    `json:"category"`
		Modified       string `json:"modified"`
		Id             string `json:"id"`
		FileId         int    `json:"fileid"`
		IsFolder       bool   `json:"isfolder"`
		ParentFolderId int    `json:"parentfolderid"`
		Size           int    `json:"size"`
		ContentType    string `json:"contenttype"`
	} `json:"metadata"`
}
