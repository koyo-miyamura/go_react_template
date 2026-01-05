package repository

import (
	"backend/internal/domain"
	"backend/internal/infra/ent"
	"context"

	"github.com/samber/lo"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	records, err := r.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	users := lo.Map(records, func(record *ent.User, _ int) domain.User {
		return domain.User{
			ID:    uint64(record.ID),
			Name:  record.Name,
			Email: record.Email,
		}
	})

	return users, nil
}
