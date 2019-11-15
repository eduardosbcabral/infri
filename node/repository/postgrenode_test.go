package repository_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduardosbcabral/infri/models"
	nodeRepo "github.com/eduardosbcabral/infri/node/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgreGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "ip", "name"}).
		AddRow(1, "192.168.0.1", "LB")

	query := "SELECT id, ip, name FROM node WHERE id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := nodeRepo.NewPostgreNodeRepository(db)

	num := int64(5)
	node, err := a.GetById(context.TODO(), num)

	assert.NoError(t, err)
	assert.NotNil(t, node)
}

func TestPostgreStore(t *testing.T) {
	node := &models.Node{
		Ip: "127.0.0.1",
		Name: "CentOS",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	query := "INSERT node SET ip=\\?, name=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(node.Ip, node.Name).
		WillReturnResult(sqlmock.NewResult(10, 1))

	n := nodeRepo.NewPostgreNodeRepository(db)

	err = n.Store(context.TODO(), node)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), node.Id)
}