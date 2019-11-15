package repository

import (
	"context"
	"database/sql"
	"github.com/eduardosbcabral/infri/models"
	"github.com/eduardosbcabral/infri/node"
	"log"
)

type postgreNodeRepository struct {
	Conn *sql.DB
}

func NewPostgreNodeRepository(Conn *sql.DB) node.Repository {
	return &postgreNodeRepository{Conn}
}

func (p *postgreNodeRepository) GetById(ctx context.Context, id int64) (*models.Node, error) {
	query := "SELECT id, ip, name FROM node WHERE id = ?"
	rows, err := p.Conn.QueryContext(ctx, query, id)
	if err != nil {
		log.Fatalf("an error ocurred when querying get by id: %s", err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	result := new(models.Node)

	for rows.Next() {
		err = rows.Scan(
			&result.Id,
			&result.Ip,
			&result.Name,
		)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return result, nil
}

func (p *postgreNodeRepository) Store(ctx context.Context, node *models.Node) error {
	query := "INSERT node SET ip=?, name=?"

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return err
	}

	res, err := stmt.ExecContext(ctx, node.Ip, node.Name)
	if err != nil {
		log.Fatal(err)
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}

	node.Id = lastId
	return nil
}