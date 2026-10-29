package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"unotek-backend/internal/delivery/http/dto"
	"unotek-backend/internal/domain"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/pagination"
	"unotek-backend/pkg/response"
)

type PartnerHandler struct {
	partnerUsecase domain.PartnerUsecase
}

func NewPartnerHandler(partnerUsecase domain.PartnerUsecase) *PartnerHandler {
	return &PartnerHandler{partnerUsecase: partnerUsecase}
}

// ListPublic godoc
// @Summary      Get list of active product & tech partners (Public)
// @Description  Retrieve partners list for ecosystem / credentials display
// @Tags         Partners
// @Produce      json
// @Param        page query int false "Page number" default(1)
// @Param        per_page query int false "Items per page" default(10)
// @Param        search query string false "Search keyword"
// @Success      200 {object} response.APIResponse{data=[]domain.ProductPartner,meta=pagination.Meta} "List of partners"
// @Router       /api/v1/partners [get]
func (h *PartnerHandler) ListPublic(c *gin.Context) {
	params := pagination.GetPaginationParams(c)
	partners, total, err := h.partnerUsecase.List(c.Request.Context(), params, true)
	if err != nil {
		response.Error(c, err)
		return
	}

	meta := pagination.BuildMeta(params, total)
	response.SuccessWithPagination(c, http.StatusOK, "Product partners retrieved successfully", partners, meta)
}

// GetPublicBySlug godoc
// @Summary      Get partner detail by slug (Public)
// @Description  Retrieve specific partner profile by unique slug
// @Tags         Partners
// @Produce      json
// @Param        slug path string true "Partner slug"
// @Success      200 {object} response.APIResponse{data=domain.ProductPartner} "Partner Detail"
// @Failure      404 {object} response.APIResponse "Partner Not Found"
// @Router       /api/v1/partners/{slug} [get]
func (h *PartnerHandler) GetPublicBySlug(c *gin.Context) {
	slug := c.Param("slug")
	partner, err := h.partnerUsecase.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Partner retrieved successfully", partner)
}

// ListAdmin godoc
// @Summary      Get all partners with filters (Admin)
// @Description  Retrieve full list of partners for admin dashboard
// @Tags         Admin - Partners
// @Produce      json
// @Security     BearerAuth
// @Param        page query int false "Page number" default(1)
// @Param        per_page query int false "Items per page" default(10)
// @Param        search query string false "Search keyword"
// @Param        sort_by query string false "Sort field" default(created_at)
// @Param        order query string false "Sort order (asc/desc)" default(desc)
// @Success      200 {object} response.APIResponse{data=[]domain.ProductPartner,meta=pagination.Meta} "List of partners"
// @Failure      401 {object} response.APIResponse "Unauthorized"
// @Router       /api/v1/admin/partners [get]
func (h *PartnerHandler) ListAdmin(c *gin.Context) {
	params := pagination.GetPaginationParams(c)
	partners, total, err := h.partnerUsecase.List(c.Request.Context(), params, false)
	if err != nil {
		response.Error(c, err)
		return
	}

	meta := pagination.BuildMeta(params, total)
	response.SuccessWithPagination(c, http.StatusOK, "Partners retrieved successfully", partners, meta)
}

// GetByID godoc
// @Summary      Get partner by ID (Admin)
// @Description  Retrieve full partner entity by UUID
// @Tags         Admin - Partners
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Partner UUID"
// @Success      200 {object} response.APIResponse{data=domain.ProductPartner} "Partner Detail"
// @Failure      404 {object} response.APIResponse "Partner Not Found"
// @Router       /api/v1/admin/partners/{id} [get]
func (h *PartnerHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid partner ID"))
		return
	}

	partner, err := h.partnerUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Partner retrieved successfully", partner)
}

// Create godoc
// @Summary      Create a new partner (Admin)
// @Description  Add new hardware or tech partner to master list
// @Tags         Admin - Partners
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreatePartnerRequest true "Partner Data"
// @Success      201 {object} response.APIResponse{data=domain.ProductPartner} "Partner Created"
// @Failure      400 {object} response.APIResponse "Validation Error"
// @Failure      409 {object} response.APIResponse "Slug Conflict"
// @Router       /api/v1/admin/partners [post]
func (h *PartnerHandler) Create(c *gin.Context) {
	var req dto.CreatePartnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	partner := domain.ProductPartner{
		Name:         req.Name,
		Slug:         req.Slug,
		PartnerType:  req.PartnerType,
		WebsiteURL:   req.WebsiteURL,
		LogoURL:      req.LogoURL,
		Description:  req.Description,
		DisplayOrder: req.DisplayOrder,
		IsActive:     *req.IsActive,
	}

	if err := h.partnerUsecase.Create(c.Request.Context(), &partner); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusCreated, "Partner created successfully", partner)
}

// Update godoc
// @Summary      Update an existing partner (Admin)
// @Description  Modify partner details by UUID
// @Tags         Admin - Partners
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Partner UUID"
// @Param        request body dto.UpdatePartnerRequest true "Updated Partner Data"
// @Success      200 {object} response.APIResponse{data=domain.ProductPartner} "Partner Updated"
// @Failure      400 {object} response.APIResponse "Validation Error"
// @Failure      404 {object} response.APIResponse "Partner Not Found"
// @Router       /api/v1/admin/partners/{id} [put]
func (h *PartnerHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid partner ID"))
		return
	}

	var req dto.UpdatePartnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	partner := domain.ProductPartner{
		ID:           id,
		Name:         req.Name,
		Slug:         req.Slug,
		PartnerType:  req.PartnerType,
		WebsiteURL:   req.WebsiteURL,
		LogoURL:      req.LogoURL,
		Description:  req.Description,
		DisplayOrder: req.DisplayOrder,
		IsActive:     *req.IsActive,
	}

	if err := h.partnerUsecase.Update(c.Request.Context(), &partner); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Partner updated successfully", partner)
}

// Delete godoc
// @Summary      Delete a partner (Admin)
// @Description  Soft delete partner by UUID
// @Tags         Admin - Partners
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Partner UUID"
// @Success      200 {object} response.APIResponse "Partner Deleted"
// @Failure      404 {object} response.APIResponse "Partner Not Found"
// @Router       /api/v1/admin/partners/{id} [delete]
func (h *PartnerHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid partner ID"))
		return
	}

	if err := h.partnerUsecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Partner deleted successfully", nil)
}
