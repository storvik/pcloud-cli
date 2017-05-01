package models

type GetfileResponse struct {
	Path    string   `json:"path"`
	Size    int      `json:"size"`
	Expires string   `json:"expires"`
	Hosts   []string `json:"hosts"`
}
