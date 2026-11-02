package usecase

import (
	"context"

	"github.com/Oleg-amur/github-release-notifier/services/release-tracker/internal/modules/releasetracker/domain"
)

type GetRepository struct {
	repositories repositoryStore
}

func NewGetRepository(repositories repositoryStore) *GetRepository {
	return &GetRepository{repositories: repositories}
}

func (u *GetRepository) Execute(ctx context.Context, id int64) (*domain.Repository, error) {
	return u.repositories.GetByID(ctx, id)
}
