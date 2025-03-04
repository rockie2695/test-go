package models

// HTTPError example
type HTTPError struct {
	// Code    int    `json:"code" example:"400"`D
	Message string `json:"message" example:"status bad request"`
}
type HTTPResponse struct {
	// Code    int    `json:"code" example:"400"`D
	Message string `json:"message" example:"ok"`
}
