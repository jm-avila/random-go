package status

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmavila/golang/web-server-4/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/ping", h.handlePing).Methods("GET")
}

func (h *Handler) handlePing(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, nil)
}
