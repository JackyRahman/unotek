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

type ClientHandler struct {
	clientUsecase domain.ClientUsecase
}

func NewClientHandler(clientUsecase domain.ClientUsecase) *ClientHandler {
	return &ClientHandler{clientUsecase: clientUsecase}
}

// ListPublic godoc
// @Summary      Get list of active clients (Public)
// @Description  Retrieve public client list with logos for landing page
// @Tags         Clients
// @Produce      json
// @Param        page query int false "Page number" default(1)
// @Param        per_page query int false "Items per page" default(10)
// @Param        search query string false "Search keyword"
// @Success      200 {object} response.APIResponse{data=[]domain.Client,meta=pagination.Meta} "List of clients"
// @Router       /api/v1/clients [get]
func (h *ClientHandler) ListPublic(c *gin.Context) {
	params := pagination.GetPaginationParams(c)
	clients, total, err := h.clientUsecase.List(c.Request.Context(), params, true)
	if err != nil {
		response.Error(c, err)
		return
	}

	meta := pagination.BuildMeta(params, total)
	response.SuccessWithPagination(c, http.StatusOK, "Clients retrieved successfully", clients, meta)
}

// GetPublicBySlug godoc
// @Summary      Get client detail by slug (Public)
// @Description  Retrieve specific client profile by unique slug
// @Tags         Clients
// @Produce      json
// @Param        slug path string true "Client slug"
// @Success      200 {object} response.APIResponse{data=domain.Client} "Client Detail"
// @Failure      404 {object} response.APIResponse "Client Not Found"
// @Router       /api/v1/clients/{slug} [get]
func (h *ClientHandler) GetPublicBySlug(c *gin.Context) {
	slug := c.Param("slug")
	client, err := h.clientUsecase.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Client retrieved successfully", client)
}

// ListAdmin godoc
// @Summary      Get all clients with filters (Admin)
// @Description  Retrieve clients list including inactive for admin dashboard
// @Tags         Admin - Clients
// @Produce      json
// @Security     BearerAuth
// @Param        page query int false "Page number" default(1)
// @Param        per_page query int false "Items per page" default(10)
// @Param        search query string false "Search keyword"
// @Param        sort_by query string false "Sort field" default(created_at)
// @Param        order query string false "Sort order (asc/desc)" default(desc)
// @Success      200 {object} response.APIResponse{data=[]domain.Client,meta=pagination.Meta} "List of clients"
// @Failure      401 {object} response.APIResponse "Unauthorized"
// @Failure      403 {object} response.APIResponse "Forbidden"
// @Router       /api/v1/admin/clients [get]
func (h *ClientHandler) ListAdmin(c *gin.Context) {
	params := pagination.GetPaginationParams(c)
	clients, total, err := h.clientUsecase.List(c.Request.Context(), params, false)
	if err != nil {
		response.Error(c, err)
		return
	}

	meta := pagination.BuildMeta(params, total)
	response.SuccessWithPagination(c, http.StatusOK, "Clients retrieved successfully", clients, meta)
}

// GetByID godoc
// @Summary      Get client detail by ID (Admin)
// @Description  Retrieve full client entity by UUID
// @Tags         Admin - Clients
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Client UUID"
// @Success      200 {object} response.APIResponse{data=domain.Client} "Client Detail"
// @Failure      404 {object} response.APIResponse "Client Not Found"
// @Router       /api/v1/admin/clients/{id} [get]
func (h *ClientHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid client ID"))
		return
	}

	client, err := h.clientUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Client retrieved successfully", client)
}

// Create godoc
// @Summary      Create a new client (Admin)
// @Description  Add new client company to master catalog
// @Tags         Admin - Clients
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateClientRequest true "Client Data"
// @Success      201 {object} response.APIResponse{data=domain.Client} "Client Created"
// @Failure      400 {object} response.APIResponse "Validation Error"
// @Failure      409 {object} response.APIResponse "Slug Conflict"
// @Router       /api/v1/admin/clients [post]
func (h *ClientHandler) Create(c *gin.Context) {
	var req dto.CreateClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	client := domain.Client{
		Name:         req.Name,
		Slug:         req.Slug,
		Industry:     req.Industry,
		WebsiteURL:   req.WebsiteURL,
		LogoURL:      req.LogoURL,
		Description:  req.Description,
		IsFeatured:   req.IsFeatured,
		DisplayOrder: req.DisplayOrder,
		IsActive:     *req.IsActive,
	}

	if err := h.clientUsecase.Create(c.Request.Context(), &client); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusCreated, "Client created successfully", client)
}

// Update godoc
// @Summary      Update an existing client (Admin)
// @Description  Modify client details by UUID
// @Tags         Admin - Clients
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Client UUID"
// @Param        request body dto.UpdateClientRequest true "Updated Client Data"
// @Success      200 {object} response.APIResponse{data=domain.Client} "Client Updated"
// @Failure      400 {object} response.APIResponse "Validation Error"
// @Failure      404 {object} response.APIResponse "Client Not Found"
// @Router       /api/v1/admin/clients/{id} [put]
func (h *ClientHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid client ID"))
		return
	}

	var req dto.UpdateClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	client := domain.Client{
		ID:           id,
		Name:         req.Name,
		Slug:         req.Slug,
		Industry:     req.Industry,
		WebsiteURL:   req.WebsiteURL,
		LogoURL:      req.LogoURL,
		Description:  req.Description,
		IsFeatured:   req.IsFeatured,
		DisplayOrder: req.DisplayOrder,
		IsActive:     *req.IsActive,
	}

	if err := h.clientUsecase.Update(c.Request.Context(), &client); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Client updated successfully", client)
}

// Delete godoc
// @Summary      Delete a client (Admin)
// @Description  Soft delete client record by UUID
// @Tags         Admin - Clients
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Client UUID"
// @Success      200 {object} response.APIResponse "Client Deleted"
// @Failure      404 {object} response.APIResponse "Client Not Found"
// @Router       /api/v1/admin/clients/{id} [delete]
func (h *ClientHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid client ID"))
		return
	}

	if err := h.clientUsecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Client deleted successfully", nil)
}
