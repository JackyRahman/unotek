package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"unotek-backend/internal/config"
	"unotek-backend/internal/repository/postgres"
	"unotek-backend/pkg/database"
)

func main() {
	var actionFlag string
	flag.StringVar(&actionFlag, "action", "", "Migration action: up, seed, all, drop")
	flag.Parse()

	action := actionFlag
	if action == "" {
		if len(os.Args) > 1 && os.Args[1] != "" {
			action = os.Args[1]
		} else {
			action = "up"
		}
	}

	cfg := config.LoadConfig()
	log.Println("==================================================")
	log.Printf("[MIGRATE CLI] Connecting to: %s\n", cfg.App.Name)
	log.Println("==================================================")

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("[DATABASE] Connection error: %v\n", err)
	}

	switch action {
	case "up", "migrate":
		log.Println("[MIGRATE CLI] Running schema migration...")
		if err := database.MigrateTables(db); err != nil {
			log.Fatalf("[MIGRATION ERROR] %v\n", err)
		}
		log.Println("[MIGRATE CLI] Schema migration completed successfully!")

	case "seed":
		log.Println("[MIGRATE CLI] Running initial seeder...")
		if err := postgres.SeedInitialData(db); err != nil {
			log.Fatalf("[SEEDER ERROR] %v\n", err)
		}
		log.Println("[MIGRATE CLI] Initial seeder completed successfully!")

	case "all":
		log.Println("[MIGRATE CLI] Running migration and seeding...")
		if err := database.MigrateTables(db); err != nil {
			log.Fatalf("[MIGRATION ERROR] %v\n", err)
		}
		if err := postgres.SeedInitialData(db); err != nil {
			log.Fatalf("[SEEDER ERROR] %v\n", err)
		}
		log.Println("[MIGRATE CLI] Migration and seeding completed successfully!")

	case "drop", "down":
		log.Println("[MIGRATE CLI] Dropping tables...")
		if err := database.DropTables(db); err != nil {
			log.Fatalf("[DROP ERROR] %v\n", err)
		}
		log.Println("[MIGRATE CLI] All tables dropped successfully!")

	default:
		fmt.Printf("Unknown action '%s'. Available actions: up, seed, all, drop\n", action)
		fmt.Println("Example:")
		fmt.Println("  go run cmd/migrate/main.go up     # Run schema migration")
		fmt.Println("  go run cmd/migrate/main.go seed   # Seed master data & super admin")
		fmt.Println("  go run cmd/migrate/main.go all    # Run both migration and seed")
		fmt.Println("  go run cmd/migrate/main.go drop   # Drop all domain tables")
	}
}
