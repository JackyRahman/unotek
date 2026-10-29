package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"unotek-backend/internal/config"
	"unotek-backend/internal/domain"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SSLMode,
		cfg.Database.Timezone,
	)

	gormLogLevel := logger.Warn
	if cfg.App.Debug {
		gormLogLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from gorm: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)

	if cfg.App.EnableAutoMigrate {
		log.Println("[DATABASE] Running automatic schema migration...")
		// Ensure pgcrypto extension for gen_random_uuid()
		db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
		db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`)

		if err := db.AutoMigrate(
			&domain.Role{},
			&domain.Permission{},
			&domain.User{},
			&domain.CompanyProfile{},
			&domain.OfficeLocation{},
			&domain.Client{},
			&domain.ProductPartner{},
			&domain.Solution{},
			&domain.SolutionFeature{},
			&domain.SolutionModule{},
			&domain.TeamMember{},
			&domain.ContactSubmission{},
		); err != nil {
			return nil, fmt.Errorf("auto migration failed: %w", err)
		}
		log.Println("[DATABASE] Schema migration completed successfully.")
	}

	return db, nil
}
