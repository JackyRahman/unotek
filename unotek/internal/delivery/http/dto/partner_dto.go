package dto

type CreatePartnerRequest struct {
	Name         string `json:"name" binding:"required" example:"Zebra Technologies"`
	Slug         string `json:"slug" example:"zebra-technologies"`
	PartnerType  string `json:"partner_type" binding:"required" example:"hardware"` // hardware, technology, cloud, iot_automation
	WebsiteURL   string `json:"website_url" example:"https://zebra.com"`
	LogoURL      string `json:"logo_url" binding:"required" example:"https://assets.company.com/partners/zebra.svg"`
	Description  string `json:"description" example:"Penyedia barcode scanner dan printer enterprise."`
	DisplayOrder int    `json:"display_order" example:"1"`
	IsActive     *bool  `json:"is_active" binding:"required" example:"true"`
}

type UpdatePartnerRequest struct {
	Name         string `json:"name" binding:"required" example:"Zebra Technologies"`
	Slug         string `json:"slug" example:"zebra-technologies"`
	PartnerType  string `json:"partner_type" binding:"required" example:"hardware"`
	WebsiteURL   string `json:"website_url" example:"https://zebra.com"`
	LogoURL      string `json:"logo_url" binding:"required" example:"https://assets.company.com/partners/zebra.svg"`
	Description  string `json:"description" example:"Penyedia barcode scanner dan printer enterprise."`
	DisplayOrder int    `json:"display_order" example:"1"`
	IsActive     *bool  `json:"is_active" binding:"required" example:"true"`
}
