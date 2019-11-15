package repository

import (
	"context"
	"github.com/eduardosbcabral/infri/models"
	"github.com/eduardosbcabral/infri/node"
	"github.com/hashicorp/go-memdb"
)

type inMemNodeRepository struct {
	DB *memdb.MemDB
}

func NewInMemNodeRepository(DB *memdb.MemDB) node.Repository {
	return &inMemNodeRepository{DB}
}

func (i *inMemNodeRepository) GetById(ctx context.Context, id int64) (*models.Node, error) {
	txn := i.DB.Txn(false)
	defer txn.Abort()

	data, err := txn.First("nodes", "id", id)
	if err != nil {
		return nil, err
	}

	node, _ := data.(*models.Node)

	return node, nil
}

func (i *inMemNodeRepository) Store(ctx context.Context, node *models.Node) error {
	txn := i.DB.Txn(true)

	err := txn.Insert("nodes", node)
	if err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func NodeTableSchema() (schema *memdb.TableSchema) {
	schema = &memdb.TableSchema{
		Name: "nodes",
		Indexes: map[string]*memdb.IndexSchema{
			"id": &memdb.IndexSchema{
				Name: "id",
				Unique: true,
				Indexer: &memdb.IntFieldIndex{Field: "Id"},
			},
			"ip": &memdb.IndexSchema{
				Name: "ip",
				Unique: false,
				Indexer: &memdb.StringFieldIndex{Field: "Ip"},
			},
			"name": &memdb.IndexSchema{
				Name: "name",
				Unique: false,
				Indexer: &memdb.StringFieldIndex{Field: "Name"},
			},
		},
	}

	return
}