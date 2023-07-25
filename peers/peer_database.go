package peers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gtfintechlab/scatter-protocol/utils"
)

func getTrainersByTopic(node *utils.PeerNode, topicName string) []string {
	rows, err := node.DataStore.Query(`
		SELECT node_id from peer_topic_maps
		WHERE (topic_name = $1) AND (node_type = $2)
	`, topicName, utils.PEER_TRAINER)

	if err != nil {
		log.Fatal(err)
	}

	trainers := []string{}
	for rows.Next() {
		var nodeId string
		rows.Scan(&nodeId)

		trainers = append(trainers, nodeId)
	}

	return trainers
}

func addTopicFromInfo(node *utils.PeerNode, nodeId string, topicName string, nodeType string, dataPath *string) {
	statement, err := node.DataStore.Prepare(`
		INSERT INTO peer_topic_maps (node_id, topic_name, node_type, data_path)
		VALUES ($1, $2, $3, $4);
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	if dataPath != nil {
		statement.Exec(nodeId, topicName, nodeType, *dataPath)
	} else {
		statement.Exec(nodeId, topicName, nodeType, nil)
	}

}

func getTopicsByNodeId(node *utils.PeerNode, nodeId string) []string {
	rows, err := node.DataStore.Query(`
		SELECT topic_name FROM peer_topic_maps
		WHERE (node_id = $1)
	`, nodeId)

	if err != nil {
		log.Fatal(err)
	}

	topics := []string{}
	for rows.Next() {
		var topicName string
		rows.Scan(&topicName)

		topics = append(topics, topicName)
	}

	return topics
}
func getTrainersForAllTopics(node *utils.PeerNode) map[string][]string {
	topicTrainerMap := map[string][]string{}
	rows, err := node.DataStore.Query(`
		SELECT node_id, node_type, topic_name 
		FROM peer_topic_maps
		WHERE (node_type = $1)`,
		utils.PEER_TRAINER,
	)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var nodeId string
		var nodeType string
		var topicName string
		rows.Scan(&nodeId, &nodeType, &topicName)
		topicTrainerMap[topicName] = append(topicTrainerMap[topicName], nodeId)
	}

	return topicTrainerMap

}

func checkIfTopicExistsForNode(node *utils.PeerNode, topicName string) bool {
	rows, err := node.DataStore.Query(`
		SELECT * 
		FROM peer_topic_maps
		WHERE (node_id = $1)`, node.NodeId.String(),
	)
	if err != nil {
		log.Fatal(err)
	}

	return rows.Next()

}

func getInitialTopics(rootDir string, node *utils.PeerNode) error {
	if node.PeerType == utils.PEER_REQUESTOR {
		return nil
	}
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != rootDir {
			// Get the depth of the current directory by counting path separators
			depth := strings.Count(path[len(rootDir):], string(os.PathSeparator))
			if depth == 1 {
				statement, err := node.DataStore.Prepare(`
					INSERT INTO peer_topic_maps (node_id, topic_name, node_type, data_path) 
					VALUES($1, $2, $3, $4);`,
				)
				if err != nil {
					log.Fatal(err)
				}
				defer statement.Close()
				statement.Exec(node.NodeId.String(), info.Name(), node.PeerType, path)
			} else if depth > 1 {
				// If the depth is greater than 1, don't traverse further into subdirectories
				return filepath.SkipDir
			}
		}
		return nil
	})

	return err
}

func connectToPostgres(peerType string) *sql.DB {
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
	connStr := fmt.Sprintf("user=postgres password=postgres host=localhost port=%d dbname=postgres sslmode=disable", dbPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Peer Node: Error connecting to PostgreSQL: %v", err)
	}

	fmt.Println("Peer Node: Connected to PostgreSQL server successfully!")
	return db
}
