package service

import (
	"github.com/Zhiyenbek/interview-result-service/config"
	"github.com/Zhiyenbek/interview-result-service/internal/models"
	"github.com/Zhiyenbek/interview-result-service/internal/repository"
	"go.uber.org/zap"
)

type ResultService interface {
	CreateResult(req *models.CreateResultRequest) error
	GetResult(publicID string) (*models.InterviewResults, error)
}

type Service struct {
	ResultService
}

func New(repos *repository.Repository, log *zap.SugaredLogger, cfg *config.Configs) *Service {
	return &Service{
		ResultService: NewResultService(repos, cfg, log),
	}
}
