package peerDatabase

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gtfintechlab/scatter-protocol/utils"
	_ "github.com/lib/pq" // PostgreSQL driver, replace with your preferred driver
)

var currentDir, _ = os.Getwd()
var (
	dbDriver      = "postgres" // Replace with your preferred database driver name
	dbUser        = "postgres"
	dbPassword    = "postgres"
	dbHost        = "localhost"
	dbName        = "postgres"
	migrationsDir = fmt.Sprintf("%s/peers/db/migrations", currentDir)
)

func MigratePeerDB(direction string, peerType string) {
	if direction != "up" && direction != "down" {
		fmt.Println("Invalid direction. Use 'up' or 'down'.")
		os.Exit(1)
	}

	db, err := connectToDatabase(peerType)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		os.Exit(1)
	}
	defer db.Close()

	migrations, err := getMigrations(direction)
	if err != nil {
		fmt.Println("Error getting migrations:", err)
		os.Exit(1)
	}

	for _, migration := range migrations {
		fmt.Printf("Running %s migration: %s\n", direction, migration.Name)
		if err := runMigration(db, migration.Path, direction); err != nil {
			fmt.Println("Error executing migration:", err)
			os.Exit(1)
		}
	}
}

func connectToDatabase(peerType string) (*sql.DB, error) {
	var cmd *exec.Cmd
	var dbPort int
	if peerType == utils.PEER_REQUESTOR {
		dbPort = 8701
		cmd = exec.Command("docker", "run", "--name", "requestor-postgres", "-e", "POSTGRES_USER=postgres", "-e", "POSTGRES_PASSWORD=postgres", "-p", "8701:5432", "-d", "postgres")
	} else {
		dbPort = 8702
		cmd = exec.Command("docker", "run", "--name", "trainer-postgres", "-e", "POSTGRES_USER=postgres", "-e", "POSTGRES_PASSWORD=postgres", "-p", "8702:5432", "-d", "postgres")
	}
	cmd.Output()

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDriver, connStr)
	if err != nil {
		return nil, err
	}
	time.Sleep(10 * time.Second)
	return db, nil
}

type migrationFile struct {
	Name string
	Path string
}

func getMigrations(direction string) ([]migrationFile, error) {
	var migrations []migrationFile

	files, err := ioutil.ReadDir(migrationsDir)
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
	content, err := ioutil.ReadFile(path)
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
