package dto

type SubmitContactRequest struct {
	FullName        string `json:"full_name" binding:"required" example:"Budi Pratama"`
	Email           string `json:"email" binding:"required,email" example:"budi@perusahaan.com"`
	Phone           string `json:"phone" example:"+628123456789"`
	CompanyName     string `json:"company_name" example:"PT Maju Mundur Logistik"`
	ServiceInterest string `json:"service_interest" example:"Warehouse Management System (WMS)"`
	Subject         string `json:"subject" binding:"required" example:"Permintaan Demo Produk WMS"`
	Message         string `json:"message" binding:"required" example:"Halo tim Unotek, kami tertarik untuk demo produk WMS untuk 3 gudang kami."`
}

type UpdateContactStatusRequest struct {
	Status string `json:"status" binding:"required" example:"contacted"` // new, in_review, contacted, resolved, archived
	Notes  string `json:"notes" example:"Sudah dihubungi oleh tim sales pada 2026-10-29"`
}
