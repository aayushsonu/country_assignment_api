package main

// swagger:ignore
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// swagger:ignore
type ErrorResponse struct {
	Error string `json:"error" swaggerignore:"true"`
}
