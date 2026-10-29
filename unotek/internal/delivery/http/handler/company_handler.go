package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"unotek-backend/internal/delivery/http/dto"
	"unotek-backend/internal/domain"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/response"
)

type CompanyHandler struct {
	companyUsecase domain.CompanyUsecase
}

func NewCompanyHandler(companyUsecase domain.CompanyUsecase) *CompanyHandler {
	return &CompanyHandler{companyUsecase: companyUsecase}
}

// GetPublicProfile godoc
// @Summary      Get company profile & brand information (Public)
// @Description  Retrieve company tagline, vision, mission, and about content
// @Tags         Company Info
// @Produce      json
// @Success      200 {object} response.APIResponse{data=domain.CompanyProfile} "Company Profile"
// @Router       /api/v1/company/profile [get]
func (h *CompanyHandler) GetPublicProfile(c *gin.Context) {
	profile, err := h.companyUsecase.GetProfile(c.Request.Context())
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Company profile retrieved successfully", profile)
}

// UpdateProfile godoc
// @Summary      Update company profile information (Admin)
// @Description  Modify company branding, vision, and mission
// @Tags         Admin - Company
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.UpdateCompanyProfileRequest true "Company Profile Data"
// @Success      200 {object} response.APIResponse{data=domain.CompanyProfile} "Profile Updated"
// @Failure      400 {object} response.APIResponse "Validation Error"
// @Router       /api/v1/admin/company/profile [put]
func (h *CompanyHandler) UpdateProfile(c *gin.Context) {
	var req dto.UpdateCompanyProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	profile := domain.CompanyProfile{
		CompanyName:     req.CompanyName,
		Tagline:         req.Tagline,
		AboutText:       req.AboutText,
		Vision:          req.Vision,
		Mission:         req.Mission,
		FoundedYear:     req.FoundedYear,
		LogoURL:         req.LogoURL,
		FaviconURL:      req.FaviconURL,
		MetaTitle:       req.MetaTitle,
		MetaDescription: req.MetaDescription,
	}

	if err := h.companyUsecase.UpdateProfile(c.Request.Context(), &profile); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Company profile updated successfully", profile)
}

// ListPublicOffices godoc
// @Summary      Get list of active office locations (Public)
// @Description  Retrieve head office & branches addresses for footer and contact page
// @Tags         Company Info
// @Produce      json
// @Success      200 {object} response.APIResponse{data=[]domain.OfficeLocation} "List of office locations"
// @Router       /api/v1/company/offices [get]
func (h *CompanyHandler) ListPublicOffices(c *gin.Context) {
	offices, err := h.companyUsecase.ListOffices(c.Request.Context(), true)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Office locations retrieved successfully", offices)
}

// ListAdminOffices godoc
// @Summary      Get all office locations (Admin)
// @Description  Retrieve all office branches including inactive
// @Tags         Admin - Company
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=[]domain.OfficeLocation} "List of office locations"
// @Router       /api/v1/admin/company/offices [get]
func (h *CompanyHandler) ListAdminOffices(c *gin.Context) {
	offices, err := h.companyUsecase.ListOffices(c.Request.Context(), false)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Office locations retrieved successfully", offices)
}

// CreateOffice godoc
// @Summary      Add new office location (Admin)
// @Description  Create office branch entry
// @Tags         Admin - Company
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateOfficeRequest true "Office Data"
// @Success      201 {object} response.APIResponse{data=domain.OfficeLocation} "Office Created"
// @Router       /api/v1/admin/company/offices [post]
func (h *CompanyHandler) CreateOffice(c *gin.Context) {
	var req dto.CreateOfficeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	office := domain.OfficeLocation{
		BranchName:    req.BranchName,
		BranchType:    req.BranchType,
		Address:       req.Address,
		City:          req.City,
		Province:      req.Province,
		PostalCode:    req.PostalCode,
		Country:       req.Country,
		Phone:         req.Phone,
		Email:         req.Email,
		WorkingHours:  req.WorkingHours,
		GoogleMapsURL: req.GoogleMapsURL,
		Latitude:      req.Latitude,
		Longitude:     req.Longitude,
		IsHeadOffice:  req.IsHeadOffice,
		DisplayOrder:  req.DisplayOrder,
		IsActive:      *req.IsActive,
	}

	if err := h.companyUsecase.CreateOffice(c.Request.Context(), &office); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusCreated, "Office location created successfully", office)
}

// UpdateOffice godoc
// @Summary      Update office location (Admin)
// @Description  Modify office branch details by UUID
// @Tags         Admin - Company
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Office UUID"
// @Param        request body dto.UpdateOfficeRequest true "Updated Office Data"
// @Success      200 {object} response.APIResponse{data=domain.OfficeLocation} "Office Updated"
// @Router       /api/v1/admin/company/offices/{id} [put]
func (h *CompanyHandler) UpdateOffice(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid office ID"))
		return
	}

	var req dto.UpdateOfficeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	office := domain.OfficeLocation{
		ID:            id,
		BranchName:    req.BranchName,
		BranchType:    req.BranchType,
		Address:       req.Address,
		City:          req.City,
		Province:      req.Province,
		PostalCode:    req.PostalCode,
		Country:       req.Country,
		Phone:         req.Phone,
		Email:         req.Email,
		WorkingHours:  req.WorkingHours,
		GoogleMapsURL: req.GoogleMapsURL,
		Latitude:      req.Latitude,
		Longitude:     req.Longitude,
		IsHeadOffice:  req.IsHeadOffice,
		DisplayOrder:  req.DisplayOrder,
		IsActive:      *req.IsActive,
	}

	if err := h.companyUsecase.UpdateOffice(c.Request.Context(), &office); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Office location updated successfully", office)
}

// DeleteOffice godoc
// @Summary      Delete office location (Admin)
// @Description  Soft delete office branch by UUID
// @Tags         Admin - Company
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Office UUID"
// @Success      200 {object} response.APIResponse "Office Deleted"
// @Router       /api/v1/admin/company/offices/{id} [delete]
func (h *CompanyHandler) DeleteOffice(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid office ID"))
		return
	}

	if err := h.companyUsecase.DeleteOffice(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Office location deleted successfully", nil)
}
