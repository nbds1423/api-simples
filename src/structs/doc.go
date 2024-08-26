package structs

type Status struct {
	// Code http
	// in: uint16
	Code uint16 `json:"code"`
	// Message
	// in: string
	Message string `json:"message"`
} //@name ErrorResponse
