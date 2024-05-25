package api

import (
	v1 "anivibe-service/internal/infrastructure/routers/http/api/v1"
	"anivibe-service/internal/infrastructure/services/feedback"
	"log"

	"github.com/gorilla/mux"
)

const prefixV1 = "/api/v1"

func SetupAPIRouters(mainRouter *mux.Router, adminUserID int64) {
	log.Println("setup routers")

	mainRouterV1 := mainRouter.PathPrefix(prefixV1).Subrouter()

	feedbackService := feedback.NewFeedbackService(adminUserID)

	v1.SetupV1MangaRouters(mainRouterV1)
	v1.SetupV1ImageRouters(mainRouterV1)
	v1.SetupV1SitemapRouters(mainRouterV1)
	v1.SetupV1FeedbackRouters(mainRouterV1, feedbackService)
}
