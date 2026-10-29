package dto

type UpdateCompanyProfileRequest struct {
	CompanyName     string `json:"company_name" binding:"required" example:"Unotek Digital Solusindo"`
	Tagline         string `json:"tagline" example:"Empowering Enterprise Logistics & Supply Chain"`
	AboutText       string `json:"about_text" example:"Kami adalah software house spesialis enterprise..."`
	Vision          string `json:"vision" example:"Menjadi mitra transformasi digital nomor satu..."`
	Mission         string `json:"mission" example:"Mengembangkan produk perangkat lunak handal..."`
	FoundedYear     int    `json:"founded_year" example:"2018"`
	LogoURL         string `json:"logo_url" example:"https://assets.company.com/brand/unotek-logo.svg"`
	FaviconURL      string `json:"favicon_url" example:"https://assets.company.com/brand/favicon.ico"`
	MetaTitle       string `json:"meta_title" example:"Unotek - Enterprise Software Solutions"`
	MetaDescription string `json:"meta_description" example:"Partner software house terpercaya..."`
}

type CreateOfficeRequest struct {
	BranchName    string   `json:"branch_name" binding:"required" example:"Jakarta Head Office"`
	BranchType    string   `json:"branch_type" binding:"required" example:"head_office"`
	Address       string   `json:"address" binding:"required" example:"Menara Standard Chartered Lt. 28"`
	City          string   `json:"city" binding:"required" example:"Jakarta Selatan"`
	Province      string   `json:"province" example:"DKI Jakarta"`
	PostalCode    string   `json:"postal_code" example:"12930"`
	Country       string   `json:"country" example:"Indonesia"`
	Phone         string   `json:"phone" example:"+62 21 5550 1928"`
	Email         string   `json:"email" example:"contact@company.com"`
	WorkingHours  string   `json:"working_hours" example:"Senin - Jumat: 08:30 - 17:30 WIB"`
	GoogleMapsURL string   `json:"google_maps_url" example:"https://maps.google.com/..."`
	Latitude      *float64 `json:"latitude"`
	Longitude     *float64 `json:"longitude"`
	IsHeadOffice  bool     `json:"is_head_office" example:"true"`
	DisplayOrder  int      `json:"display_order" example:"1"`
	IsActive      *bool    `json:"is_active" binding:"required" example:"true"`
}

type UpdateOfficeRequest struct {
	BranchName    string   `json:"branch_name" binding:"required" example:"Jakarta Head Office"`
	BranchType    string   `json:"branch_type" binding:"required" example:"head_office"`
	Address       string   `json:"address" binding:"required" example:"Menara Standard Chartered Lt. 28"`
	City          string   `json:"city" binding:"required" example:"Jakarta Selatan"`
	Province      string   `json:"province" example:"DKI Jakarta"`
	PostalCode    string   `json:"postal_code" example:"12930"`
	Country       string   `json:"country" example:"Indonesia"`
	Phone         string   `json:"phone" example:"+62 21 5550 1928"`
	Email         string   `json:"email" example:"contact@company.com"`
	WorkingHours  string   `json:"working_hours" example:"Senin - Jumat: 08:30 - 17:30 WIB"`
	GoogleMapsURL string   `json:"google_maps_url" example:"https://maps.google.com/..."`
	Latitude      *float64 `json:"latitude"`
	Longitude     *float64 `json:"longitude"`
	IsHeadOffice  bool     `json:"is_head_office" example:"true"`
	DisplayOrder  int      `json:"display_order" example:"1"`
	IsActive      *bool    `json:"is_active" binding:"required" example:"true"`
}
