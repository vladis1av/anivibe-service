package feedback

import (
	"anivibe-service/internal/infrastructure/services/feedback"
	"encoding/json"
	"fmt"
	"net/http"
)

type FeedbackHandler struct {
	feedbackService *feedback.FeedbackService
}

func NewFeedbackHandler(feedbackService *feedback.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{
		feedbackService: feedbackService,
	}
}

func (h *FeedbackHandler) SendFeedback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var feedbackReq feedback.FeedbackRequest
	if err := json.NewDecoder(r.Body).Decode(&feedbackReq); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := h.feedbackService.SendFeedback(&feedbackReq); err != nil {
		http.Error(w, fmt.Sprintf("Failed to send feedback: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Feedback sent successfully"))
}
