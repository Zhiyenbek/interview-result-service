package repository

import (
	"github.com/Zhiyenbek/interview-result-service/config"
	"github.com/Zhiyenbek/interview-result-service/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type ResultRepository interface {
	CreateResult(result *models.InterviewResults) error
	GetResult(publicID string) (*models.InterviewResults, error)
}
type Repository struct {
	ResultRepository
}

func New(db *pgxpool.Pool, cfg *config.Configs, log *zap.SugaredLogger) *Repository {
	return &Repository{
		ResultRepository: NewResultRepository(db, cfg.DB, log),
	}
}
