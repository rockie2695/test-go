package models

// HTTPError example
type HTTPError struct {
	// Code    int    `json:"code" example:"400"`D
	Message string `json:"message" example:"status bad request"`
}