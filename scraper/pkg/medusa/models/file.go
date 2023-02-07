package models

type FileUploadResponse struct {
	Uploads []struct {
		URL string `json:"url"`
	} `json:"uploads"`
}
