package v1

import (
	"anivibe-service/internal/infrastructure/handlers/http/v1/feedback"
	feedbackService "anivibe-service/internal/infrastructure/services/feedback"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupV1FeedbackRouters(router *mux.Router, service *feedbackService.FeedbackService) {
	log.Println("setup feedback routers v1")

	feedbackHandler := feedback.NewFeedbackHandler(service)

	feedback := router.PathPrefix("/feedback").Subrouter()

	feedback.HandleFunc("", feedbackHandler.SendFeedback).Methods(http.MethodPost)
}
