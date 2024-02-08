package service

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

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

func (s *resultService) GetResult(publicID string) (*models.InterviewResults, error) {
	res, err := s.resultRepo.GetResult(publicID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *resultService) CreateResult(req *models.CreateResultRequest) error {
	result := &models.InterviewResults{}
	for i, question := range req.Questions {
		for j, res := range question.EmotionResults {
			path := strings.Split(question.Video, "/")
			length := len(path)
			prefix := res.Emotion
			newPath := ""
			for i, val := range path {
				if i == length-1 {
					err := os.Mkdir(newPath+val+"_emotionResults", os.ModePerm)
					if err != nil && !errors.Is(err, os.ErrExist) {
						return err
					}
					newPath = newPath + val + "_emotionResults" + "/" + uuid.NewString() + "_" + prefix + "_" + val
				} else {
					newPath = newPath + val + "/"
				}
			}
			err := ffmpeg.Input(question.Video, ffmpeg.KwArgs{"ss": int(res.ExactTime)}).
				Output(newPath, ffmpeg.KwArgs{"t": int(res.Duration)}).OverWriteOutput().Run()
			if err != nil {
				s.logger.Error(err)
				continue
			}
			req.Questions[i].EmotionResults[j].VideoPath = newPath
		}
	}
	var err error
	result.Result, err = json.Marshal(req)
	if err != nil {
		return err
	}
	return s.resultRepo.CreateResult(result)
}
