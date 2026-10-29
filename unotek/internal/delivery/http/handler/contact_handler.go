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

type ContactHandler struct {
	contactUsecase domain.ContactUsecase
}

func NewContactHandler(contactUsecase domain.ContactUsecase) *ContactHandler {
	return &ContactHandler{contactUsecase: contactUsecase}
}

// SubmitPublic godoc
// @Summary      Submit a contact us message / lead (Public)
// @Description  Submit customer inquiry or consultation request
// @Tags         Contact Us
// @Accept       json
// @Produce      json
// @Param        request body dto.SubmitContactRequest true "Contact Form Data"
// @Success      201 {object} response.APIResponse "Inquiry Submitted Successfully"
// @Failure      400 {object} response.APIResponse "Validation Error"
// @Router       /api/v1/contact [post]
func (h *ContactHandler) SubmitPublic(c *gin.Context) {
	var req dto.SubmitContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	submission := domain.ContactSubmission{
		FullName:        req.FullName,
		Email:           req.Email,
		Phone:           req.Phone,
		CompanyName:     req.CompanyName,
		ServiceInterest: req.ServiceInterest,
		Subject:         req.Subject,
		Message:         req.Message,
		IPAddress:       c.ClientIP(),
		UserAgent:       c.Request.UserAgent(),
	}

	if err := h.contactUsecase.Submit(c.Request.Context(), &submission); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusCreated, "Thank you. Your message has been received and our team will contact you soon.", nil)
}

// ListAdmin godoc
// @Summary      Get submitted contact messages / leads (Admin)
// @Description  Retrieve submitted contact requests with status filter
// @Tags         Admin - Contacts
// @Produce      json
// @Security     BearerAuth
// @Param        page query int false "Page number" default(1)
// @Param        per_page query int false "Items per page" default(10)
// @Param        status query string false "Status filter (new, in_review, contacted, resolved, archived)"
// @Param        search query string false "Search keyword"
// @Param        sort_by query string false "Sort field" default(created_at)
// @Param        order query string false "Sort order (asc/desc)" default(desc)
// @Success      200 {object} response.APIResponse{data=[]domain.ContactSubmission,meta=pagination.Meta} "List of inquiries"
// @Failure      401 {object} response.APIResponse "Unauthorized"
// @Router       /api/v1/admin/contacts [get]
func (h *ContactHandler) ListAdmin(c *gin.Context) {
	params := pagination.GetPaginationParams(c)
	statusFilter := c.Query("status")
	submissions, total, err := h.contactUsecase.List(c.Request.Context(), params, statusFilter)
	if err != nil {
		response.Error(c, err)
		return
	}

	meta := pagination.BuildMeta(params, total)
	response.SuccessWithPagination(c, http.StatusOK, "Contact submissions retrieved successfully", submissions, meta)
}

// GetByID godoc
// @Summary      Get contact message detail by ID (Admin)
// @Description  Retrieve detailed inquiry record
// @Tags         Admin - Contacts
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Contact UUID"
// @Success      200 {object} response.APIResponse{data=domain.ContactSubmission} "Contact Detail"
// @Failure      404 {object} response.APIResponse "Contact Not Found"
// @Router       /api/v1/admin/contacts/{id} [get]
func (h *ContactHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid contact ID"))
		return
	}

	submission, err := h.contactUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Contact submission retrieved successfully", submission)
}

// UpdateStatus godoc
// @Summary      Update inquiry status & notes (Admin)
// @Description  Change follow-up status (e.g. new -> contacted -> resolved)
// @Tags         Admin - Contacts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Contact UUID"
// @Param        request body dto.UpdateContactStatusRequest true "Status update data"
// @Success      200 {object} response.APIResponse "Status Updated"
// @Failure      400 {object} response.APIResponse "Validation Error"
// @Failure      404 {object} response.APIResponse "Contact Not Found"
// @Router       /api/v1/admin/contacts/{id}/status [patch]
func (h *ContactHandler) UpdateStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid contact ID"))
		return
	}

	var req dto.UpdateContactStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	if err := h.contactUsecase.UpdateStatus(c.Request.Context(), id, req.Status, req.Notes); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Contact status updated successfully", nil)
}

// Delete godoc
// @Summary      Delete a contact inquiry (Admin)
// @Description  Soft delete contact submission by UUID
// @Tags         Admin - Contacts
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Contact UUID"
// @Success      200 {object} response.APIResponse "Contact Deleted"
// @Failure      404 {object} response.APIResponse "Contact Not Found"
// @Router       /api/v1/admin/contacts/{id} [delete]
func (h *ContactHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid contact ID"))
		return
	}

	if err := h.contactUsecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Contact submission deleted successfully", nil)
}
