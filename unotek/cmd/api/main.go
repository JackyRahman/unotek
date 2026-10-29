package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"unotek-backend/internal/config"
	"unotek-backend/internal/delivery/http/handler"
	"unotek-backend/internal/delivery/http/router"
	"unotek-backend/internal/repository/postgres"
	"unotek-backend/internal/usecase"
	"unotek-backend/pkg/database"
	"unotek-backend/pkg/security"
)

// @title           Unotek Enterprise Backend API
// @version         1.0.0
// @description     Professional Go Clean Architecture Backend with PostgreSQL, JWT RBAC, Swagger, and Content Management.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  dev@company.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your JWT token.
func main() {
	// 1. Load Configurations
	cfg := config.LoadConfig()
	log.Printf("[APP] Starting %s in %s mode...\n", cfg.App.Name, cfg.App.Env)

	// 2. Initialize Database Connection
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("[DATABASE] Connection error: %v\n", err)
	}

	// 3. Run Automatic Initial Seeder if enabled
	if cfg.App.EnableAutoSeeder {
		if err := postgres.SeedInitialData(db); err != nil {
			log.Printf("[SEEDER] Warning during seeding: %v\n", err)
		}
	}

	// 4. Initialize Core Security Services
	jwtService := security.NewJWTService(cfg.JWT.Secret, cfg.JWT.ExpirationHours)

	// 5. Initialize Repositories (PostgreSQL)
	rbacRepo := postgres.NewRBACRepository(db)
	userRepo := postgres.NewUserRepository(db)
	clientRepo := postgres.NewClientRepository(db)
	partnerRepo := postgres.NewPartnerRepository(db)
	solutionRepo := postgres.NewSolutionRepository(db)
	teamRepo := postgres.NewTeamRepository(db)
	contactRepo := postgres.NewContactRepository(db)
	companyRepo := postgres.NewCompanyRepository(db)

	// 6. Initialize Usecases (Business Logic)
	authUsecase := usecase.NewAuthUsecase(userRepo, rbacRepo, jwtService)
	rbacUsecase := usecase.NewRBACUsecase(rbacRepo)
	clientUsecase := usecase.NewClientUsecase(clientRepo)
	partnerUsecase := usecase.NewPartnerUsecase(partnerRepo)
	solutionUsecase := usecase.NewSolutionUsecase(solutionRepo)
	teamUsecase := usecase.NewTeamUsecase(teamRepo)
	contactUsecase := usecase.NewContactUsecase(contactRepo)
	companyUsecase := usecase.NewCompanyUsecase(companyRepo)

	// 7. Initialize HTTP Handlers
	handlers := &router.Handlers{
		AuthHandler:     handler.NewAuthHandler(authUsecase),
		ClientHandler:   handler.NewClientHandler(clientUsecase),
		PartnerHandler:  handler.NewPartnerHandler(partnerUsecase),
		SolutionHandler: handler.NewSolutionHandler(solutionUsecase),
		TeamHandler:     handler.NewTeamHandler(teamUsecase),
		ContactHandler:  handler.NewContactHandler(contactUsecase),
		CompanyHandler:  handler.NewCompanyHandler(companyUsecase),
		RBACHandler:     handler.NewRBACHandler(rbacUsecase),
	}

	// 8. Setup Gin Router
	r := router.SetupRouter(cfg, jwtService, handlers)

	// 9. Start HTTP Server with Graceful Shutdown
	serverAddr := fmt.Sprintf(":%s", cfg.App.Port)
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("[SERVER] Listening and serving HTTP on port %s...\n", cfg.App.Port)
		log.Printf("[SWAGGER] Interactive API Documentation: http://localhost:%s/swagger/index.html\n", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[SERVER] Failed to start server: %v\n", err)
		}
	}()

	// Graceful Shutdown on SIGINT or SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("[SERVER] Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[SERVER] Server forced to shutdown: %v\n", err)
	}

	log.Println("[SERVER] Server exited successfully.")
}
