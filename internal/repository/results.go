package repository

import (
	"context"

	"github.com/Zhiyenbek/interview-result-service/config"
	"github.com/Zhiyenbek/interview-result-service/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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
		INSERT INTO interviews (public_id, results)
		VALUES ($1, $2);
	`

	_, err = r.db.Exec(ctx, query, uuid.NewString(), result.Result)
	if err != nil {
		r.logger.Errorf("Error occurred while inserting interview result: %v", err)
		return err
	}

	return nil
}

func (r *resultRepository) GetResult(publicID string) (*models.InterviewResults, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()

	query := `
		SELECT results
		FROM interviews
		WHERE public_id = $1;
	`

	var resultBytes []byte

	err := r.db.QueryRow(ctx, query, publicID).Scan(&resultBytes)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // Interview result not found
		}

		r.logger.Errorf("Error occurred while retrieving interview result: %v", err)
		return nil, err
	}

	interviewResult := &models.InterviewResults{
		PublicID: publicID,
		Result:   resultBytes,
	}

	return interviewResult, nil
}
