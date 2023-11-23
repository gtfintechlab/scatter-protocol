package peerDatabase

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq" // PostgreSQL driver, replace with your preferred driver
)

var currentDir, _ = os.Getwd()
var (
	migrationsDir = fmt.Sprintf("%s/peers/db/migrations", currentDir)
)

func MigratePeerDB(direction string, peerType string, databaseUsername string, databasePassword string, databasePort int) {
	if direction != "up" && direction != "down" {
		log.Fatal("Invalid direction. Use 'up' or 'down'.")
		os.Exit(1)
	}

	db := connectToPostgres(peerType, databaseUsername, databasePassword, databasePort)
	defer db.Close()

	migrations, err := getMigrations(direction)
	if err != nil {
		log.Fatal("Error getting migrations:", err)
		os.Exit(1)
	}

	for _, migration := range migrations {
		log.Printf("Running %s migration: %s\n", direction, migration.Name)
		if err := runMigration(db, migration.Path, direction); err != nil {
			log.Fatal("Error executing migration:", err)
			os.Exit(1)
		}
	}
}

func connectToPostgres(peerType string, username string, password string, port int) *sql.DB {
	connStr := fmt.Sprintf(
		"user=%s password=%s host=localhost port=%d dbname=postgres sslmode=disable",
		username, password, port,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Peer Node: Error connecting to PostgreSQL: %v", err)
	}

	log.Println("Peer Node: Connected to PostgreSQL server successfully!")
	return db
}

type migrationFile struct {
	Name string
	Path string
}

func getMigrations(direction string) ([]migrationFile, error) {
	var migrations []migrationFile

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			migration := migrationFile{
				Name: strings.TrimSuffix(file.Name(), ".sql"),
				Path: filepath.Join(migrationsDir, file.Name()),
			}
			migrations = append(migrations, migration)
		}
	}

	if direction == "down" {
		// Reverse the order for "down" migrations
		for i, j := 0, len(migrations)-1; i < j; i, j = i+1, j-1 {
			migrations[i], migrations[j] = migrations[j], migrations[i]
		}
	}

	return migrations, nil
}

func runMigration(db *sql.DB, path string, direction string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	upMigration := strings.Split(string(content), "-- Down")[0]
	downMigration := strings.Split(string(content), "-- Down")[1]

	if direction == "up" {
		_, err = db.Exec(string(upMigration))
		if err != nil {
			return err
		}
	} else {
		_, err = db.Exec(string(downMigration))
		if err != nil {
			return err
		}

	}

	return nil
}
