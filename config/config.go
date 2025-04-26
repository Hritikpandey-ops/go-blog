package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the specified .env file
func LoadEnv(fileName string) {
	err := godotenv.Load(fileName)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// GetDBConfig retrieves the database configuration from environment variables
func GetDBConfig() (string, string, string, string, string) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	return dbUser, dbPassword, dbName, dbHost, dbPort
}

// RunMigrations runs the database migrations
func RunMigrations() {
	// Load environment variables
	LoadEnv(".env")

	// Get database configuration
	dbUser, dbPassword, dbName, dbHost, dbPort := GetDBConfig()

	// Build the PostgreSQL connection URL
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Get absolute path for migrations
	migrationsPath, err := filepath.Abs("./migrations")
	if err != nil {
		log.Fatalf("Error getting absolute path for migrations: %v", err)
	}
	fmt.Println("Migration path:", migrationsPath)

	// Convert Windows path to URL format
	// Method 1: Using filepath.ToSlash
	migrationsURL := "file://" + filepath.ToSlash(migrationsPath)

	// OR Method 2: More robust URI encoding
	// migrationsURL := "file:///" + strings.ReplaceAll(filepath.ToSlash(migrationsPath), " ", "%20")

	// Debug output
	fmt.Println("Migration URL:", migrationsURL)

	// Check if migration file exists
	migrationFile := filepath.Join(migrationsPath, "001_initial_migration.up.sql")
	if _, err := os.Stat(migrationFile); os.IsNotExist(err) {
		log.Fatalf("Migration file does not exist: %v", err)
	}

	// Set up migration
	m, err := migrate.New(migrationsURL, dbURL)
	if err != nil {
		log.Fatalf("Migration setup error: %v", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error applying migrations: %v", err)
	}

	fmt.Println("Migrations applied successfully")
}
