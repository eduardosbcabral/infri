package node

import (
	"context"
	"github.com/eduardosbcabral/infri/models"
)

type UseCase interface {
	GetById(ctx context.Context, id int64) (*models.Node, error)
	Save(ctx context.Context, node *models.Node) error
}