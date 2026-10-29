package dto

type CreateClientRequest struct {
	Name         string `json:"name" binding:"required" example:"PT Indomarco Prismatama"`
	Slug         string `json:"slug" example:"indomaret"`
	Industry     string `json:"industry" example:"Retail & FMCG"`
	WebsiteURL   string `json:"website_url" example:"https://indomaret.co.id"`
	LogoURL      string `json:"logo_url" binding:"required" example:"https://assets.company.com/clients/indomaret.svg"`
	Description  string `json:"description" example:"Implementasi sistem monitoring distribusi toko."`
	IsFeatured   bool   `json:"is_featured" example:"true"`
	DisplayOrder int    `json:"display_order" example:"1"`
	IsActive     *bool  `json:"is_active" binding:"required" example:"true"`
}

type UpdateClientRequest struct {
	Name         string `json:"name" binding:"required" example:"PT Indomarco Prismatama"`
	Slug         string `json:"slug" example:"indomaret"`
	Industry     string `json:"industry" example:"Retail & FMCG"`
	WebsiteURL   string `json:"website_url" example:"https://indomaret.co.id"`
	LogoURL      string `json:"logo_url" binding:"required" example:"https://assets.company.com/clients/indomaret.svg"`
	Description  string `json:"description" example:"Implementasi sistem monitoring distribusi toko."`
	IsFeatured   bool   `json:"is_featured" example:"true"`
	DisplayOrder int    `json:"display_order" example:"1"`
	IsActive     *bool  `json:"is_active" binding:"required" example:"true"`
}
