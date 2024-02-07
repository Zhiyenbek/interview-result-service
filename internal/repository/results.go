package repository

import (
	"context"

	"github.com/Zhiyenbek/interview-result-service/config"
	"github.com/Zhiyenbek/interview-result-service/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

type resultRepository struct {
	db     *pgxpool.Pool
	cfg    *config.DBConf
	logger *zap.SugaredLogger
}

func NewResultRepository(db *pgxpool.Pool, cfg *config.DBConf, logger *zap.SugaredLogger) ResultRepository {
	return &resultRepository{
		db:     db,
		cfg:    cfg,
		logger: logger,
	}
}

func (r *resultRepository) CreateResult(result *models.InterviewResults) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	var err error
	query := `
		INSERT INTO interview_results (public_id, result)
		VALUES ($1, $2);
	`

	_, err = r.db.Exec(ctx, query, result.PublicID, pq.Array(result.Result))
	if err != nil {
		r.logger.Errorf("Error occurred while inserting interview result: %v", err)
		return err
	}

	return nil
}
