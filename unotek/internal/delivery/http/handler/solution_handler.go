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

type SolutionHandler struct {
	solutionUsecase domain.SolutionUsecase
}

func NewSolutionHandler(solutionUsecase domain.SolutionUsecase) *SolutionHandler {
	return &SolutionHandler{solutionUsecase: solutionUsecase}
}

// ListPublic godoc
// @Summary      Get list of solutions (Public)
// @Description  Retrieve public catalog of software solutions (e.g. WMS, TMS)
// @Tags         Solutions
// @Produce      json
// @Param        page query int false "Page number" default(1)
// @Param        per_page query int false "Items per page" default(10)
// @Param        search query string false "Search keyword"
// @Success      200 {object} response.APIResponse{data=[]domain.Solution,meta=pagination.Meta} "List of solutions"
// @Router       /api/v1/solutions [get]
func (h *SolutionHandler) ListPublic(c *gin.Context) {
	params := pagination.GetPaginationParams(c)
	solutions, total, err := h.solutionUsecase.List(c.Request.Context(), params, true)
	if err != nil {
		response.Error(c, err)
		return
	}

	meta := pagination.BuildMeta(params, total)
	response.SuccessWithPagination(c, http.StatusOK, "Solutions retrieved successfully", solutions, meta)
}

// GetPublicBySlug godoc
// @Summary      Get solution detail by slug (Public)
// @Description  Retrieve solution including features and modules by slug (e.g. warehouse-management-system)
// @Tags         Solutions
// @Produce      json
// @Param        slug path string true "Solution slug"
// @Success      200 {object} response.APIResponse{data=domain.Solution} "Solution Detail"
// @Failure      404 {object} response.APIResponse "Solution Not Found"
// @Router       /api/v1/solutions/{slug} [get]
func (h *SolutionHandler) GetPublicBySlug(c *gin.Context) {
	slug := c.Param("slug")
	solution, err := h.solutionUsecase.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Solution retrieved successfully", solution)
}

// ListAdmin godoc
// @Summary      Get all solutions (Admin)
// @Description  Retrieve all solutions for CMS dashboard
// @Tags         Admin - Solutions
// @Produce      json
// @Security     BearerAuth
// @Param        page query int false "Page number" default(1)
// @Param        per_page query int false "Items per page" default(10)
// @Param        search query string false "Search keyword"
// @Param        sort_by query string false "Sort field" default(created_at)
// @Param        order query string false "Sort order (asc/desc)" default(desc)
// @Success      200 {object} response.APIResponse{data=[]domain.Solution,meta=pagination.Meta} "List of solutions"
// @Failure      401 {object} response.APIResponse "Unauthorized"
// @Router       /api/v1/admin/solutions [get]
func (h *SolutionHandler) ListAdmin(c *gin.Context) {
	params := pagination.GetPaginationParams(c)
	solutions, total, err := h.solutionUsecase.List(c.Request.Context(), params, false)
	if err != nil {
		response.Error(c, err)
		return
	}

	meta := pagination.BuildMeta(params, total)
	response.SuccessWithPagination(c, http.StatusOK, "Solutions retrieved successfully", solutions, meta)
}

// GetByID godoc
// @Summary      Get solution detail by ID (Admin)
// @Description  Retrieve full solution entity by UUID
// @Tags         Admin - Solutions
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Solution UUID"
// @Success      200 {object} response.APIResponse{data=domain.Solution} "Solution Detail"
// @Failure      404 {object} response.APIResponse "Solution Not Found"
// @Router       /api/v1/admin/solutions/{id} [get]
func (h *SolutionHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid solution ID"))
		return
	}

	solution, err := h.solutionUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Solution retrieved successfully", solution)
}

// Create godoc
// @Summary      Create a new solution (Admin)
// @Description  Add new solution catalog entry
// @Tags         Admin - Solutions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateSolutionRequest true "Solution Data"
// @Success      201 {object} response.APIResponse{data=domain.Solution} "Solution Created"
// @Failure      400 {object} response.APIResponse "Validation Error"
// @Failure      409 {object} response.APIResponse "Slug Conflict"
// @Router       /api/v1/admin/solutions [post]
func (h *SolutionHandler) Create(c *gin.Context) {
	var req dto.CreateSolutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	solution := domain.Solution{
		Title:        req.Title,
		Slug:         req.Slug,
		Tagline:      req.Tagline,
		Summary:      req.Summary,
		Description:  req.Description,
		IconURL:      req.IconURL,
		HeroImageURL: req.HeroImageURL,
		BrochureURL:  req.BrochureURL,
		DisplayOrder: req.DisplayOrder,
		IsFeatured:   req.IsFeatured,
		IsActive:     *req.IsActive,
	}

	if err := h.solutionUsecase.Create(c.Request.Context(), &solution); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusCreated, "Solution created successfully", solution)
}

// Update godoc
// @Summary      Update an existing solution (Admin)
// @Description  Modify solution entity by UUID
// @Tags         Admin - Solutions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Solution UUID"
// @Param        request body dto.UpdateSolutionRequest true "Updated Solution Data"
// @Success      200 {object} response.APIResponse{data=domain.Solution} "Solution Updated"
// @Failure      400 {object} response.APIResponse "Validation Error"
// @Failure      404 {object} response.APIResponse "Solution Not Found"
// @Router       /api/v1/admin/solutions/{id} [put]
func (h *SolutionHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid solution ID"))
		return
	}

	var req dto.UpdateSolutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	solution := domain.Solution{
		ID:           id,
		Title:        req.Title,
		Slug:         req.Slug,
		Tagline:      req.Tagline,
		Summary:      req.Summary,
		Description:  req.Description,
		IconURL:      req.IconURL,
		HeroImageURL: req.HeroImageURL,
		BrochureURL:  req.BrochureURL,
		DisplayOrder: req.DisplayOrder,
		IsFeatured:   req.IsFeatured,
		IsActive:     *req.IsActive,
	}

	if err := h.solutionUsecase.Update(c.Request.Context(), &solution); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Solution updated successfully", solution)
}

// Delete godoc
// @Summary      Delete a solution (Admin)
// @Description  Soft delete solution by UUID
// @Tags         Admin - Solutions
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Solution UUID"
// @Success      200 {object} response.APIResponse "Solution Deleted"
// @Failure      404 {object} response.APIResponse "Solution Not Found"
// @Router       /api/v1/admin/solutions/{id} [delete]
func (h *SolutionHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid solution ID"))
		return
	}

	if err := h.solutionUsecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Solution deleted successfully", nil)
}

// AddFeature godoc
// @Summary      Add feature item to solution (Admin)
// @Tags         Admin - Solutions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Solution UUID"
// @Param        request body dto.CreateFeatureRequest true "Feature Data"
// @Success      201 {object} response.APIResponse{data=domain.SolutionFeature} "Feature Added"
// @Router       /api/v1/admin/solutions/{id}/features [post]
func (h *SolutionHandler) AddFeature(c *gin.Context) {
	idStr := c.Param("id")
	solutionID, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid solution ID"))
		return
	}

	var req dto.CreateFeatureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	feature := domain.SolutionFeature{
		SolutionID:   solutionID,
		Title:        req.Title,
		Description:  req.Description,
		IconURL:      req.IconURL,
		DisplayOrder: req.DisplayOrder,
		IsActive:     *req.IsActive,
	}

	if err := h.solutionUsecase.AddFeature(c.Request.Context(), &feature); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusCreated, "Feature added successfully", feature)
}

// DeleteFeature godoc
// @Summary      Delete feature item from solution (Admin)
// @Tags         Admin - Solutions
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Solution UUID"
// @Param        feature_id path string true "Feature UUID"
// @Success      200 {object} response.APIResponse "Feature Deleted"
// @Router       /api/v1/admin/solutions/{id}/features/{feature_id} [delete]
func (h *SolutionHandler) DeleteFeature(c *gin.Context) {
	featureIDStr := c.Param("feature_id")
	featureID, err := uuid.Parse(featureIDStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid feature ID"))
		return
	}

	if err := h.solutionUsecase.DeleteFeature(c.Request.Context(), featureID); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Feature deleted successfully", nil)
}

// AddModule godoc
// @Summary      Add module item to solution (Admin)
// @Tags         Admin - Solutions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Solution UUID"
// @Param        request body dto.CreateModuleRequest true "Module Data"
// @Success      201 {object} response.APIResponse{data=domain.SolutionModule} "Module Added"
// @Router       /api/v1/admin/solutions/{id}/modules [post]
func (h *SolutionHandler) AddModule(c *gin.Context) {
	idStr := c.Param("id")
	solutionID, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid solution ID"))
		return
	}

	var req dto.CreateModuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	module := domain.SolutionModule{
		SolutionID:   solutionID,
		Title:        req.Title,
		Description:  req.Description,
		ImageURL:     req.ImageURL,
		DisplayOrder: req.DisplayOrder,
		IsActive:     *req.IsActive,
	}

	if err := h.solutionUsecase.AddModule(c.Request.Context(), &module); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusCreated, "Module added successfully", module)
}

// DeleteModule godoc
// @Summary      Delete module item from solution (Admin)
// @Tags         Admin - Solutions
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Solution UUID"
// @Param        module_id path string true "Module UUID"
// @Success      200 {object} response.APIResponse "Module Deleted"
// @Router       /api/v1/admin/solutions/{id}/modules/{module_id} [delete]
func (h *SolutionHandler) DeleteModule(c *gin.Context) {
	moduleIDStr := c.Param("module_id")
	moduleID, err := uuid.Parse(moduleIDStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid module ID"))
		return
	}

	if err := h.solutionUsecase.DeleteModule(c.Request.Context(), moduleID); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Module deleted successfully", nil)
}
