package health

// HealthResponse represents the health check response
type HealthResponse struct {
	Status  string `json:"status" example:"ok"`
	Message string `json:"message" example:"Vibecart server is running"`
}

// DBTestResponse represents the database connection test response
type DBTestResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Aiven connected successfully"`
}

// DBTestErrorResponse represents the database connection test error response
type DBTestErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"Aiven connection failed"`
	Error   string `json:"error" example:"connection timeout"`
}

// CloudinaryTestResponse represents the Cloudinary connection test response
type CloudinaryTestResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Cloudinary connected successfully"`
}

// CloudinaryTestErrorResponse represents the Cloudinary connection test error response
type CloudinaryTestErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"Cloudinary connection failed"`
	Error   string `json:"error" example:"invalid credentials"`
}
