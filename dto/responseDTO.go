package dto

// Generic response from server
type GenericResponseDTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
