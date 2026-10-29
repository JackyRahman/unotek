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

type TeamHandler struct {
	teamUsecase domain.TeamUsecase
}

func NewTeamHandler(teamUsecase domain.TeamUsecase) *TeamHandler {
	return &TeamHandler{teamUsecase: teamUsecase}
}

// ListPublic godoc
// @Summary      Get list of team members (Public)
// @Description  Retrieve team members and leadership for About Us page
// @Tags         Team
// @Produce      json
// @Param        page query int false "Page number" default(1)
// @Param        per_page query int false "Items per page" default(10)
// @Param        search query string false "Search keyword"
// @Success      200 {object} response.APIResponse{data=[]domain.TeamMember,meta=pagination.Meta} "List of team members"
// @Router       /api/v1/team [get]
func (h *TeamHandler) ListPublic(c *gin.Context) {
	params := pagination.GetPaginationParams(c)
	members, total, err := h.teamUsecase.List(c.Request.Context(), params, true)
	if err != nil {
		response.Error(c, err)
		return
	}

	meta := pagination.BuildMeta(params, total)
	response.SuccessWithPagination(c, http.StatusOK, "Team members retrieved successfully", members, meta)
}

// ListAdmin godoc
// @Summary      Get all team members (Admin)
// @Description  Retrieve all team members for admin dashboard
// @Tags         Admin - Team
// @Produce      json
// @Security     BearerAuth
// @Param        page query int false "Page number" default(1)
// @Param        per_page query int false "Items per page" default(10)
// @Param        search query string false "Search keyword"
// @Param        sort_by query string false "Sort field" default(created_at)
// @Param        order query string false "Sort order (asc/desc)" default(desc)
// @Success      200 {object} response.APIResponse{data=[]domain.TeamMember,meta=pagination.Meta} "List of team members"
// @Failure      401 {object} response.APIResponse "Unauthorized"
// @Router       /api/v1/admin/team [get]
func (h *TeamHandler) ListAdmin(c *gin.Context) {
	params := pagination.GetPaginationParams(c)
	members, total, err := h.teamUsecase.List(c.Request.Context(), params, false)
	if err != nil {
		response.Error(c, err)
		return
	}

	meta := pagination.BuildMeta(params, total)
	response.SuccessWithPagination(c, http.StatusOK, "Team members retrieved successfully", members, meta)
}

// GetByID godoc
// @Summary      Get team member detail by ID (Admin)
// @Description  Retrieve full team member record by UUID
// @Tags         Admin - Team
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Team Member UUID"
// @Success      200 {object} response.APIResponse{data=domain.TeamMember} "Team Member Detail"
// @Failure      404 {object} response.APIResponse "Team Member Not Found"
// @Router       /api/v1/admin/team/{id} [get]
func (h *TeamHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid team member ID"))
		return
	}

	member, err := h.teamUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Team member retrieved successfully", member)
}

// Create godoc
// @Summary      Create a new team member (Admin)
// @Description  Add team member profile
// @Tags         Admin - Team
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateTeamMemberRequest true "Team Member Data"
// @Success      201 {object} response.APIResponse{data=domain.TeamMember} "Team Member Created"
// @Failure      400 {object} response.APIResponse "Validation Error"
// @Router       /api/v1/admin/team [post]
func (h *TeamHandler) Create(c *gin.Context) {
	var req dto.CreateTeamMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	member := domain.TeamMember{
		FullName:     req.FullName,
		Position:     req.Position,
		Department:   req.Department,
		Bio:          req.Bio,
		PhotoURL:     req.PhotoURL,
		LinkedinURL:  req.LinkedinURL,
		Email:        req.Email,
		DisplayOrder: req.DisplayOrder,
		IsLeadership: req.IsLeadership,
		IsActive:     *req.IsActive,
	}

	if err := h.teamUsecase.Create(c.Request.Context(), &member); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusCreated, "Team member created successfully", member)
}

// Update godoc
// @Summary      Update an existing team member (Admin)
// @Description  Modify team member profile by UUID
// @Tags         Admin - Team
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Team Member UUID"
// @Param        request body dto.UpdateTeamMemberRequest true "Updated Team Member Data"
// @Success      200 {object} response.APIResponse{data=domain.TeamMember} "Team Member Updated"
// @Failure      400 {object} response.APIResponse "Validation Error"
// @Failure      404 {object} response.APIResponse "Team Member Not Found"
// @Router       /api/v1/admin/team/{id} [put]
func (h *TeamHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid team member ID"))
		return
	}

	var req dto.UpdateTeamMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	member := domain.TeamMember{
		ID:           id,
		FullName:     req.FullName,
		Position:     req.Position,
		Department:   req.Department,
		Bio:          req.Bio,
		PhotoURL:     req.PhotoURL,
		LinkedinURL:  req.LinkedinURL,
		Email:        req.Email,
		DisplayOrder: req.DisplayOrder,
		IsLeadership: req.IsLeadership,
		IsActive:     *req.IsActive,
	}

	if err := h.teamUsecase.Update(c.Request.Context(), &member); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Team member updated successfully", member)
}

// Delete godoc
// @Summary      Delete a team member (Admin)
// @Description  Soft delete team member record by UUID
// @Tags         Admin - Team
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Team Member UUID"
// @Success      200 {object} response.APIResponse "Team Member Deleted"
// @Failure      404 {object} response.APIResponse "Team Member Not Found"
// @Router       /api/v1/admin/team/{id} [delete]
func (h *TeamHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid team member ID"))
		return
	}

	if err := h.teamUsecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Team member deleted successfully", nil)
}
