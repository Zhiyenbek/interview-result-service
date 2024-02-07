package models

type CreateResultRequest struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	Video          string          `json:"video_link"`
	EmotionResults []EmotionResult `json:"emomtion_results"`
}

type EmotionResult struct {
	Emotion   string  `json:"emotion"`
	ExactTime float64 `json:"exact_time"`
	Duration  float64 `json:"duration"`
}

type InterviewResults struct {
	PublicID string
	Result   []byte
}
