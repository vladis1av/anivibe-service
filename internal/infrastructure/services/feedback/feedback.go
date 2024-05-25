package feedback

import (
	"anivibe-service/internal/telegram"
	"encoding/base64"
	"fmt"
)

type FeedbackService struct {
	adminUserID int64
}

type FeedbackRequest struct {
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Images  []string `json:"images"`
}

func NewFeedbackService(adminUserID int64) *FeedbackService {
	return &FeedbackService{
		adminUserID: adminUserID,
	}
}

func (s *FeedbackService) SendFeedback(req *FeedbackRequest) error {
	form := &telegram.FeedbackForm{
		Name:    req.Name,
		Email:   req.Email,
		Subject: req.Subject,
		Message: req.Message,
	}

	if len(req.Images) > 5 {
		return fmt.Errorf("cannot send more than 5 images")
	}

	var imageBytes [][]byte
	for _, img := range req.Images {
		imgBytes, err := base64.StdEncoding.DecodeString(img)
		if err != nil {
			return fmt.Errorf("failed to decode image: %w", err)
		}
		imageBytes = append(imageBytes, imgBytes)
	}

	if err := telegram.SendFeedback(s.adminUserID, form, imageBytes); err != nil {
		return fmt.Errorf("failed to send feedback: %w", err)
	}

	return nil
}
