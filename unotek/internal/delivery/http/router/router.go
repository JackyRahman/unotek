package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/time/rate"
	_ "unotek-backend/docs"
	"unotek-backend/internal/config"
	"unotek-backend/internal/delivery/http/handler"
	"unotek-backend/internal/delivery/http/middleware"
	"unotek-backend/pkg/response"
	"unotek-backend/pkg/security"
)

type Handlers struct {
	AuthHandler     *handler.AuthHandler
	ClientHandler   *handler.ClientHandler
	PartnerHandler  *handler.PartnerHandler
	SolutionHandler *handler.SolutionHandler
	TeamHandler     *handler.TeamHandler
	ContactHandler  *handler.ContactHandler
	CompanyHandler  *handler.CompanyHandler
	RBACHandler     *handler.RBACHandler
}

func SetupRouter(cfg *config.Config, jwtService security.JWTService, h *Handlers) *gin.Engine {
	if !cfg.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Global Middlewares
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.RequestIDMiddleware())
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())

	// IP Rate Limiter
	rateLimiter := middleware.NewIPRateLimiter(rate.Limit(cfg.Security.RateLimitRPS), cfg.Security.RateLimitBurst)
	r.Use(middleware.RateLimiterMiddleware(rateLimiter))

	// Swagger Documentation Route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		response.Success(c, http.StatusOK, "Service is healthy and running", gin.H{
			"app":     cfg.App.Name,
			"version": "1.0.0",
			"status":  "UP",
		})
	})

	apiV1 := r.Group("/api/v1")
	{
		// 1. Authentication
		authGroup := apiV1.Group("/auth")
		{
			authGroup.POST("/login", h.AuthHandler.Login)
			authGroup.GET("/profile", middleware.AuthMiddleware(jwtService), h.AuthHandler.GetProfile)
		}

		// 2. Public Master & Landing Page Endpoints
		apiV1.GET("/company/profile", h.CompanyHandler.GetPublicProfile)
		apiV1.GET("/company/offices", h.CompanyHandler.ListPublicOffices)

		apiV1.GET("/clients", h.ClientHandler.ListPublic)
		apiV1.GET("/clients/:slug", h.ClientHandler.GetPublicBySlug)

		apiV1.GET("/partners", h.PartnerHandler.ListPublic)
		apiV1.GET("/partners/:slug", h.PartnerHandler.GetPublicBySlug)

		apiV1.GET("/solutions", h.SolutionHandler.ListPublic)
		apiV1.GET("/solutions/:slug", h.SolutionHandler.GetPublicBySlug)

		apiV1.GET("/team", h.TeamHandler.ListPublic)

		apiV1.POST("/contact", h.ContactHandler.SubmitPublic)

		// 3. Admin & CMS Endpoints (Protected by JWT and RBAC)
		admin := apiV1.Group("/admin")
		admin.Use(middleware.AuthMiddleware(jwtService))
		{
			// Clients Management
			clients := admin.Group("/clients")
			{
				clients.GET("", middleware.RequirePermission("clients:read"), h.ClientHandler.ListAdmin)
				clients.GET("/:id", middleware.RequirePermission("clients:read"), h.ClientHandler.GetByID)
				clients.POST("", middleware.RequirePermission("clients:create"), h.ClientHandler.Create)
				clients.PUT("/:id", middleware.RequirePermission("clients:update"), h.ClientHandler.Update)
				clients.DELETE("/:id", middleware.RequirePermission("clients:delete"), h.ClientHandler.Delete)
			}

			// Partners Management
			partners := admin.Group("/partners")
			{
				partners.GET("", middleware.RequirePermission("partners:read"), h.PartnerHandler.ListAdmin)
				partners.GET("/:id", middleware.RequirePermission("partners:read"), h.PartnerHandler.GetByID)
				partners.POST("", middleware.RequirePermission("partners:create"), h.PartnerHandler.Create)
				partners.PUT("/:id", middleware.RequirePermission("partners:update"), h.PartnerHandler.Update)
				partners.DELETE("/:id", middleware.RequirePermission("partners:delete"), h.PartnerHandler.Delete)
			}

			// Solutions Management
			solutions := admin.Group("/solutions")
			{
				solutions.GET("", middleware.RequirePermission("solutions:read"), h.SolutionHandler.ListAdmin)
				solutions.GET("/:id", middleware.RequirePermission("solutions:read"), h.SolutionHandler.GetByID)
				solutions.POST("", middleware.RequirePermission("solutions:create"), h.SolutionHandler.Create)
				solutions.PUT("/:id", middleware.RequirePermission("solutions:update"), h.SolutionHandler.Update)
				solutions.DELETE("/:id", middleware.RequirePermission("solutions:delete"), h.SolutionHandler.Delete)

				// Solution Features
				solutions.POST("/:id/features", middleware.RequirePermission("solutions:update"), h.SolutionHandler.AddFeature)
				solutions.DELETE("/:id/features/:feature_id", middleware.RequirePermission("solutions:update"), h.SolutionHandler.DeleteFeature)

				// Solution Modules
				solutions.POST("/:id/modules", middleware.RequirePermission("solutions:update"), h.SolutionHandler.AddModule)
				solutions.DELETE("/:id/modules/:module_id", middleware.RequirePermission("solutions:update"), h.SolutionHandler.DeleteModule)
			}

			// Team Members Management
			team := admin.Group("/team")
			{
				team.GET("", middleware.RequirePermission("team:read"), h.TeamHandler.ListAdmin)
				team.GET("/:id", middleware.RequirePermission("team:read"), h.TeamHandler.GetByID)
				team.POST("", middleware.RequirePermission("team:create"), h.TeamHandler.Create)
				team.PUT("/:id", middleware.RequirePermission("team:update"), h.TeamHandler.Update)
				team.DELETE("/:id", middleware.RequirePermission("team:delete"), h.TeamHandler.Delete)
			}

			// Contact Inquiries / Leads Management
			contacts := admin.Group("/contacts")
			{
				contacts.GET("", middleware.RequirePermission("contacts:read"), h.ContactHandler.ListAdmin)
				contacts.GET("/:id", middleware.RequirePermission("contacts:read"), h.ContactHandler.GetByID)
				contacts.PATCH("/:id/status", middleware.RequirePermission("contacts:update"), h.ContactHandler.UpdateStatus)
				contacts.DELETE("/:id", middleware.RequirePermission("contacts:delete"), h.ContactHandler.Delete)
			}

			// Company Profile & Locations
			company := admin.Group("/company")
			{
				company.PUT("/profile", middleware.RequirePermission("company:update"), h.CompanyHandler.UpdateProfile)
				company.GET("/offices", middleware.RequirePermission("company:update"), h.CompanyHandler.ListAdminOffices)
				company.POST("/offices", middleware.RequirePermission("company:update"), h.CompanyHandler.CreateOffice)
				company.PUT("/offices/:id", middleware.RequirePermission("company:update"), h.CompanyHandler.UpdateOffice)
				company.DELETE("/offices/:id", middleware.RequirePermission("company:update"), h.CompanyHandler.DeleteOffice)
			}

			// RBAC Management (Super Admin only)
			rbac := admin.Group("/rbac")
			rbac.Use(middleware.RequireRole("super_admin"))
			{
				rbac.GET("/roles", h.RBACHandler.GetRoles)
				rbac.POST("/roles", h.RBACHandler.CreateRole)
				rbac.POST("/roles/:id/permissions", h.RBACHandler.AssignPermissions)
				rbac.GET("/permissions", h.RBACHandler.GetPermissions)
			}
		}
	}

	return r
}
