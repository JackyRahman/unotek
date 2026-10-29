package dto

type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required" example:"admin@company.com"`
	Password   string `json:"password" binding:"required" example:"admin12345"`
}
