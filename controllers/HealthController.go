package controllers

import (
	"net/http"

	"github.com/NiranjanShetty8/bookmarkapp/web"
	"github.com/gorilla/mux"
)

type HealthController struct{}

func (hc *HealthController) healthCheck(w http.ResponseWriter, r *http.Request) {
	web.RespondJSON(&w, http.StatusOK, `{"msg":"OK"}`)
}

func (hc *HealthController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/health", hc.healthCheck).Methods("GET")
}

func NewHealthController() *HealthController {
	return &HealthController{}
}
