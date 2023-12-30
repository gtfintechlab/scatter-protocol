package peerDatabase

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/gtfintechlab/scatter-protocol/core/utils"
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

func GetEvaluationJobDataFromAddressAndTopic(node *utils.PeerNode, nodeAddress string, topicName string) string {
	rows, err := node.DataStore.Query(`
		SELECT evaluation_job_data FROM topic_mappings
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

func AddTopicFromInfo(node *utils.PeerNode, nodeAddress string, ipfsCid string, topicName string, dataPath *string, evaluationJob *string, evaluationJobData *string) {
	statement, err := node.DataStore.Prepare(`
		INSERT INTO topic_mappings (node_address, ipfs_id, topic_name, data_path, evaluation_job, evaluation_job_data)
		VALUES ($1, $2, $3, $4, $5, $6);
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	if dataPath != nil {
		statement.Exec(strings.ToLower(nodeAddress), strings.ToLower(ipfsCid), topicName, *dataPath, nil, nil)
	} else {
		statement.Exec(strings.ToLower(nodeAddress), strings.ToLower(ipfsCid), topicName, nil, *evaluationJob, *evaluationJobData)
	}

}

func WriteDecryptionKey(node *utils.PeerNode, trainerAddress string, privateKey []byte, topicName string, requestorAddress string) {
	trainerAddress = strings.ToLower(trainerAddress)
	requestorAddress = strings.ToLower(requestorAddress)

	statement, err := node.DataStore.Prepare(`
		INSERT INTO decryption_keys (trainer_address, requestor_address, topic_name, private_key)
		VALUES ($1, $2, $3, $4);
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	statement.Exec(trainerAddress, requestorAddress, topicName, privateKey)

}

func RetrieveDecryptionKey(node *utils.PeerNode, trainerAddress string, topicName string, requestorAddress string) []byte {
	trainerAddress = strings.ToLower(trainerAddress)
	requestorAddress = strings.ToLower(requestorAddress)

	rows, err := node.DataStore.Query(`
		SELECT private_key FROM decryption_keys
		WHERE (trainer_address = $1) AND (topic_name = $2) AND (requestor_address = $3);
	`, trainerAddress, topicName, requestorAddress)

	if err != nil {
		log.Fatal(err)
	}

	keys := [][]byte{}
	for rows.Next() {
		var key []byte
		rows.Scan(&key)
		keys = append(keys, key)
	}

	if len(keys) == 0 {
		return []byte{}
	}

	return keys[0]
}

func ConnectToPostgres(peerType string, username string, password string, port int) *sql.DB {
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
