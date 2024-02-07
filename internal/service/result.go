package service

import (
	"github.com/Zhiyenbek/interview-result-service/config"
	"github.com/Zhiyenbek/interview-result-service/internal/models"
	"github.com/Zhiyenbek/interview-result-service/internal/repository"
	"github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
)

type resultService struct {
	cfg        *config.Configs
	logger     *zap.SugaredLogger
	resultRepo repository.ResultRepository
}

func NewResultService(repo *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger) ResultService {
	return &resultService{
		resultRepo: repo.ResultRepository,
		cfg:        cfg,
		logger:     logger,
	}
}

func (s *resultService) CreateResult(req *models.CreateResultRequest) error {
	for _, question := range req.Questions {
		for _, res := range question.EmotionResults {
			err := ffmpeg.Input(question.Video, ffmpeg.KwArgs{"ss": res.ExactTime}).
				Output(question.Video+uuid.NewString(), ffmpeg.KwArgs{"t": res.Duration}).OverWriteOutput().Run()
			if err != nil {
				return err
			}
		}
	}
	
	return nil
}
