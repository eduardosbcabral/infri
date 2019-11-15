package repository_test

import (
	"context"
	"github.com/eduardosbcabral/infri/models"
	nodeRepo "github.com/eduardosbcabral/infri/node/repository"
	"github.com/hashicorp/go-memdb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemGetById(t *testing.T) {
	db := InitializeDatabase(t)

	nodeRepository := nodeRepo.NewInMemNodeRepository(db)

	nodeToSave := &models.Node{
		Id: int64(1),
		Ip: "127.0.0.1",
		Name: "LB",
	}

	err := nodeRepository.Store(context.TODO(), nodeToSave)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when saving structs", err)
	}

	id := int64(1)
	node, err := nodeRepository.GetById(context.TODO(), id)
	assert.NoError(t, err)
	assert.NotNil(t, node)_
}

func TestInMemStore(t *testing.T) {
	db := InitializeDatabase(t)

	nodeRepository := nodeRepo.NewInMemNodeRepository(db)

	nodeToSave := &models.Node{
		Id: int64(1),
		Ip: "127.0.0.1",
		Name: "LB",
	}

	err := nodeRepository.Store(context.TODO(), nodeToSave)
	assert.NoError(t, err)
}

func InitializeDatabase(t *testing.T) (db *memdb.MemDB) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"nodes": nodeRepo.NodeTableSchema(),
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating an in memory database", err)
	}

	return
}