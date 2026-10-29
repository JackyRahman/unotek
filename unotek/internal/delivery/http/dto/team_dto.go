package dto

type CreateTeamMemberRequest struct {
	FullName     string `json:"full_name" binding:"required" example:"Arya Wibowo"`
	Position     string `json:"position" binding:"required" example:"Chief Executive Officer"`
	Department   string `json:"department" example:"Executive"`
	Bio          string `json:"bio" example:"15+ tahun pengalaman memimpin transformasi sistem digital."`
	PhotoURL     string `json:"photo_url" example:"https://assets.company.com/team/arya.jpg"`
	LinkedinURL  string `json:"linkedin_url" example:"https://linkedin.com/in/aryawibowo"`
	Email        string `json:"email" example:"arya@company.com"`
	DisplayOrder int    `json:"display_order" example:"1"`
	IsLeadership bool   `json:"is_leadership" example:"true"`
	IsActive     *bool  `json:"is_active" binding:"required" example:"true"`
}

type UpdateTeamMemberRequest struct {
	FullName     string `json:"full_name" binding:"required" example:"Arya Wibowo"`
	Position     string `json:"position" binding:"required" example:"Chief Executive Officer"`
	Department   string `json:"department" example:"Executive"`
	Bio          string `json:"bio" example:"15+ tahun pengalaman memimpin transformasi sistem digital."`
	PhotoURL     string `json:"photo_url" example:"https://assets.company.com/team/arya.jpg"`
	LinkedinURL  string `json:"linkedin_url" example:"https://linkedin.com/in/aryawibowo"`
	Email        string `json:"email" example:"arya@company.com"`
	DisplayOrder int    `json:"display_order" example:"1"`
	IsLeadership bool   `json:"is_leadership" example:"true"`
	IsActive     *bool  `json:"is_active" binding:"required" example:"true"`
}
