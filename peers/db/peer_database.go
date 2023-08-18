package peerDatabase

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/gtfintechlab/scatter-protocol/utils"
)

func GetDatapathFromAddressAndIpfs(node *utils.PeerNode, nodeAddress string, ipfsCid string) string {
	rows, err := node.DataStore.Query(`
		SELECT data_path FROM topic_mappings
		WHERE (node_address = $1) AND (ipfs_id = $2);
	`, strings.ToLower(nodeAddress), strings.ToLower(ipfsCid))

	if err != nil {
		log.Fatal(err)
	}
	dataPaths := []string{}
	for rows.Next() {
		var dataPath string
		rows.Scan(&dataPath)
		dataPaths = append(dataPaths, dataPath)
	}

	if len(dataPaths) == 0 {
		return ""
	}

	return dataPaths[0]

}

func GetEvaluationJobFromAddressAndTopic(node *utils.PeerNode, nodeAddress string, topicName string) string {
	rows, err := node.DataStore.Query(`
		SELECT evaluation_job FROM topic_mappings
		WHERE (node_address = $1) AND (topic_name = $2);
	`, strings.ToLower(nodeAddress), topicName)

	if err != nil {
		log.Fatal(err)
	}
	evalJobPaths := []string{}
	for rows.Next() {
		var jobPath string
		rows.Scan(&jobPath)
		evalJobPaths = append(evalJobPaths, jobPath)
	}

	if len(evalJobPaths) == 0 {
		return ""
	}

	return evalJobPaths[0]

}

func AddTopicFromInfo(node *utils.PeerNode, nodeAddress string, ipfsCid string, topicName string, dataPath *string, evaluationJob *string) {
	// node.DatastoreLock.Lock()
	// defer node.DatastoreLock.Unlock()
	statement, err := node.DataStore.Prepare(`
		INSERT INTO topic_mappings (node_address, ipfs_id, topic_name, data_path, evaluation_job)
		VALUES ($1, $2, $3, $4, $5);
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	if dataPath != nil {
		statement.Exec(strings.ToLower(nodeAddress), strings.ToLower(ipfsCid), topicName, *dataPath, nil)
	} else {
		statement.Exec(strings.ToLower(nodeAddress), strings.ToLower(ipfsCid), topicName, nil, *evaluationJob)
	}

}

func ConnectToPostgres(peerType string, username string, password string, port int) *sql.DB {
	exec.Command(
		"docker", "run",
		"--name", fmt.Sprintf("%s-postgres", peerType),
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
		log.Fatalf("Peer Node: Error connecting to PostgreSQL: %v", err)
	}

	log.Println("Peer Node: Connected to PostgreSQL server successfully!")
	return db
}
