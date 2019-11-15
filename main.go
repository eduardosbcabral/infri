package main

import(
	"context"
	"github.com/eduardosbcabral/infri/external_file"
	nodeRepo "github.com/eduardosbcabral/infri/node/repository"
	"github.com/hashicorp/go-memdb"
	"log"
	"net/http"
)

func main() {

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"nodes": nodeRepo.NodeTableSchema(),
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Fatalf("error when creating database application: %s", err)
	}

	nodeRepository := nodeRepo.NewInMemNodeRepository(db)

	fileContent, err := external_file.OpenFile("nodes.json")

	nodes, _ := external_file.MapNodeFromFileContent(fileContent)

	for _, node := range nodes {
		nodeRepository.Store(context.TODO(), &node)
	}

	server := &Server{nodeRepository}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
