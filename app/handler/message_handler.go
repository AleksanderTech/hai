package handler

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/oaroz/hai/app/service"
	"github.com/gorilla/mux"
)

type handler struct {
	service service.MessageService
}

func RegisterHandlers(r *mux.Router, service service.MessageService) {
	hand := handler{service}
	r.Path("/api/messages").HandlerFunc(hand.getMessages).Methods("GET")
	r.Path("/api/messages").HandlerFunc(hand.postMessage).Methods("POST")
	r.Path("/api/messages/{id}").HandlerFunc(hand.deleteMessage).Methods("DELETE")
}

func (h handler) getMessages(w http.ResponseWriter, req *http.Request) {
	email := req.URL.Query().Get("email")
	message, err := h.service.Get(email)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)

	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}

func (h handler) postMessage(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func (h handler) deleteMessage(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
