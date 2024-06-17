package http

import (
	"gitlab.com/EndowTheGreat/onionhead/transport/ws"

	"github.com/go-chi/chi"
)

type Handler struct {
	Router  *chi.Mux
	Service string
}

func NewHandler() *Handler {
	return &Handler{
		Router: chi.NewMux(),
	}
}

func (h *Handler) SetupRoutes() {
	router := h.Router
	router.HandleFunc("/ws", ws.HandleConnections)
	router.HandleFunc("/", h.Home)
	router.HandleFunc("/share", h.Share)
	go ws.HandleMessages()
}
