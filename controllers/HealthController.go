package controllers

import (
	"net/http"

	"github.com/NiranjanShetty8/bookmarkapp/web"
	"github.com/gorilla/mux"
)

type HealthController struct{}

// Simple health check
func (hc *HealthController) healthCheck(w http.ResponseWriter, r *http.Request) {
	web.RespondJSON(&w, http.StatusOK, `{"msg":"OK"}`)
}

// Register routes in router
func (hc *HealthController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/health", hc.healthCheck).Methods("GET")
}

// Returns instance of HealthController
func NewHealthController() *HealthController {
	return &HealthController{}
}
