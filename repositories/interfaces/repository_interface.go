package interfaces

import (
	"context"
)

type RepositoryInterface[T any] interface {
	Create(ctx context.Context, model *T) error
	FindByID(ctx context.Context, id uint) (*T, error)
	Update(ctx context.Context, model *T) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, page, limit int) ([]*T, error)
	Count(ctx context.Context) (int64, error)
	// WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
