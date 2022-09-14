package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handle Pointer - stores
type Hander struct {
	Router *mux.Router
}

// newHandler - returs a pointer to a Handler
func NewHandler() *Hander {
	return &Hander{}
}

func (h *Hander) HandleRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I Still Alive")
	})


}
