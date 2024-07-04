package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmavila/golang/web-server-4/models"
	"github.com/jmavila/golang/web-server-4/utils"
)

type Handler struct {
	store models.ProductStore
}

func NewHandler(store models.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, ps)
}
