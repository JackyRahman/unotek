package dto

type CreateSolutionRequest struct {
	Title        string `json:"title" binding:"required" example:"Warehouse Management System (WMS)"`
	Slug         string `json:"slug" example:"warehouse-management-system"`
	Tagline      string `json:"tagline" example:"Smart, Scalable, and Real-Time Warehouse Control"`
	Summary      string `json:"summary" example:"Sistem otomasi gudang modern."`
	Description  string `json:"description" example:"Deskripsi lengkap mengenai sistem..."`
	IconURL      string `json:"icon_url" example:"https://assets.company.com/icons/wms.svg"`
	HeroImageURL string `json:"hero_image_url" example:"https://assets.company.com/mockups/hero.png"`
	BrochureURL  string `json:"brochure_url" example:"https://assets.company.com/docs/wms-brochure.pdf"`
	DisplayOrder int    `json:"display_order" example:"1"`
	IsFeatured   bool   `json:"is_featured" example:"true"`
	IsActive     *bool  `json:"is_active" binding:"required" example:"true"`
}

type UpdateSolutionRequest struct {
	Title        string `json:"title" binding:"required" example:"Warehouse Management System (WMS)"`
	Slug         string `json:"slug" example:"warehouse-management-system"`
	Tagline      string `json:"tagline" example:"Smart, Scalable, and Real-Time Warehouse Control"`
	Summary      string `json:"summary" example:"Sistem otomasi gudang modern."`
	Description  string `json:"description" example:"Deskripsi lengkap mengenai sistem..."`
	IconURL      string `json:"icon_url" example:"https://assets.company.com/icons/wms.svg"`
	HeroImageURL string `json:"hero_image_url" example:"https://assets.company.com/mockups/hero.png"`
	BrochureURL  string `json:"brochure_url" example:"https://assets.company.com/docs/wms-brochure.pdf"`
	DisplayOrder int    `json:"display_order" example:"1"`
	IsFeatured   bool   `json:"is_featured" example:"true"`
	IsActive     *bool  `json:"is_active" binding:"required" example:"true"`
}

type CreateFeatureRequest struct {
	Title        string `json:"title" binding:"required" example:"Real-Time Inventory Visibility"`
	Description  string `json:"description" example:"Pantau kuantiti dan lokasi stok di berbagai gudang."`
	IconURL      string `json:"icon_url" example:"https://assets.company.com/icons/eye.svg"`
	DisplayOrder int    `json:"display_order" example:"1"`
	IsActive     *bool  `json:"is_active" binding:"required" example:"true"`
}

type CreateModuleRequest struct {
	Title        string `json:"title" binding:"required" example:"Inbound & Putaway Automation"`
	Description  string `json:"description" example:"Penerimaan PO barang dan validasi kuantiti."`
	ImageURL     string `json:"image_url" example:"https://assets.company.com/modules/inbound.png"`
	DisplayOrder int    `json:"display_order" example:"1"`
	IsActive     *bool  `json:"is_active" binding:"required" example:"true"`
}
