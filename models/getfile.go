package models

// GetfileResponse contains server response after getfile call
type GetfileResponse struct {
	Path    string   `json:"path"`
	Size    int      `json:"size"`
	Expires string   `json:"expires"`
	Hosts   []string `json:"hosts"`
}
