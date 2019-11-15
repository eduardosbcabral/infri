package node

import (
	"context"
	"github.com/eduardosbcabral/infri/models"
)

type Repository interface {
	GetById(ctx context.Context, id int64) (*models.Node, error)
	Store(ctx context.Context, node *models.Node) error
}