package external_file

import (
	"encoding/json"
	"github.com/eduardosbcabral/infri/models"
	"io/ioutil"
	"log"
	"os"
)

func OpenFile(fileName string) ([]byte, error) {
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		_, err := os.Create(fileName)
		if err != nil {
			log.Fatalf("error when creating a file: %s", err)
			return nil, err
		}
	}

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error when reading a file: %s", err)
		return nil, err
	}

	return fileContent, nil
}

func WriteToFile(fileName string, fileContent string) error {
	return ioutil.WriteFile(fileName, []byte(fileContent), 0644)
}

func MapNodeFromFileContent(fileContent []byte) ([]models.Node, error) {

	var node []models.Node

	err := json.Unmarshal(fileContent, &node)
	if err != nil {
		return nil, err
	}

	return node, nil
}