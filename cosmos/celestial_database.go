package cosmos

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"

	"github.com/gtfintechlab/scatter-protocol/utils"
)

func addTopicMappingFromInfo(node *utils.CelestialNode, nodeId string, nodeType string, topicName string) {
	statement, _ := node.DataStore.Prepare(`
	INSERT INTO topic_mappings (node_id, node_type, topic_name) 
	VALUES($1, $2, $3);`,
	)
	defer statement.Close()
	statement.Exec(nodeId, utils.PEER_REQUESTOR, topicName)

}

func checkIfTopicMappingExists(node *utils.CelestialNode, nodeId string, topicName string) bool {
	rows, err := node.DataStore.Query(
		`SELECT * FROM topic_mappings 
			WHERE (node_id = $1) AND (topic_name = $2);`,
		nodeId, topicName,
	)
	if err != nil {
		log.Fatal(err)
	}
	return rows.Next()
}

func getAllTopics(node *utils.CelestialNode) []map[string]interface{} {
	rows, err := node.DataStore.Query(`
	SELECT node_id, node_type, topic_name 
	FROM topic_mappings;`,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	allTopics := []map[string]interface{}{}

	for rows.Next() {
		var nodeId string
		var nodeType string
		var topicName string
		rows.Scan(&nodeId, &nodeType, &topicName)

		allTopics = append(allTopics, map[string]interface{}{
			"node_id":    nodeId,
			"node_type":  nodeType,
			"topic_name": topicName,
		})
	}
	return allTopics
}

func deleteTopicByInfo(node *utils.CelestialNode, nodeId string, topicName string) {
	statement, _ := node.DataStore.Prepare(`
	DELETE FROM topic_mappings 
	WHERE (node_id = $1) AND (topic_name = $2);`,
	)
	defer statement.Close()
	statement.Exec(nodeId, topicName)
}

func connectToPostgres(username string, password string, port int) *sql.DB {
	exec.Command(
		"docker", "run",
		"--name", "celestial-postgres",
		"-e", fmt.Sprintf("POSTGRES_USER=%s", username),
		"-e", fmt.Sprintf("POSTGRES_PASSWORD=%s", password),
		"-p", fmt.Sprintf("%d:5432", port),
		"-d", "postgres",
	).Output()

	connStr := fmt.Sprintf(
		"user=%s password=%s host=localhost port=%d dbname=postgres sslmode=disable",
		username, password, port,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Celestial Node: Error connecting to PostgreSQL: %v", err)
	}

	fmt.Println("Celestial Node: Connected to PostgreSQL server successfully!")
	return db
}
