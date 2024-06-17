package http

import (
	"embed"
	"net/http"

	"gitlab.com/EndowTheGreat/onionhead/qr"
)

//go:embed web/index.html
var webpage embed.FS

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	page, err := webpage.ReadFile("web/index.html")
	if err != nil {
		http.Error(w, "Could not load webpage", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(page)
}

func (h *Handler) Share(w http.ResponseWriter, r *http.Request) {
	code, err := qr.GenerateQR(h.Service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "image/png")
	if _, err := w.Write(code); err != nil {
		http.Error(w, "Could not display QR image", http.StatusInternalServerError)
	}
}
