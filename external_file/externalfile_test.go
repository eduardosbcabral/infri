package external_file_test

import (
	"github.com/eduardosbcabral/infri/external_file"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOpenFileAndMapToNodeStruct(t *testing.T) {
	fileName := "test_open_file_and_map_to_node_struct.json"

	jsonContent :=
`[
  {
    "id": 555,
    "ip": "127.0.0.1",
    "name": "Teste arquivo"
  }
]`

	err := external_file.WriteToFile(fileName, jsonContent)
	assert.Nil(t, err)

	fileContent, err := external_file.OpenFile(fileName)
	assert.Nil(t, err)

	node, err := external_file.MapNodeFromFileContent(fileContent)
	assert.Nil(t, err)

	assert.Equal(t, int64(555), node[0].Id)

	os.Remove(fileName)
}

func TestOpenFile(t *testing.T) {
	fileName := "test_open_file.json"

	_, err := external_file.OpenFile(fileName)

	assert.Nil(t, err)

	os.Remove(fileName)
}

func TestReadFile(t *testing.T) {
	fileName := "test_read_file.json"

	json :=
`[
  {
    "id": 555,
    "ip": "127.0.0.1",
    "name": "Teste arquivo"
  }
]`

	err := external_file.WriteToFile(fileName, json)
	if err != nil {
		t.Fatalf("an error ocurred when writing to a file: %s", err)
	}

	fileContent, err := external_file.OpenFile(fileName)
	if err != nil {
		t.Fatalf("an error ocurred when opening a file: %s", err)
	}

	assert.Equal(t, []byte(json), fileContent)

	os.Remove(fileName)
}
