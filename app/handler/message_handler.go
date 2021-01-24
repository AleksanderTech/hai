package handler

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/oaroz/hai/app/mapper"
	"bitbucket.org/oaroz/hai/app/model"

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
	r.Path("/api/messages").HandlerFunc(hand.deleteMessage).Methods("DELETE")
}

func (h handler) getMessages(w http.ResponseWriter, req *http.Request) {
	email := req.URL.Query().Get("email")
	messages := h.service.Get(email)

	json.NewEncoder(w).Encode(messages)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h handler) postMessage(w http.ResponseWriter, req *http.Request) {
	var msgReq model.CreateMessageRequest
	err := json.NewDecoder(req.Body).Decode(&msgReq)
	if err != nil {
		http.Error(w, "Cannot Deserialize body to 'Message' format", http.StatusBadRequest)
	}
	var res model.CreateMessageResponse = mapper.MessageToCreateResponse(h.service.Create(mapper.CreateReqToMessage(msgReq)))

	json.NewEncoder(w).Encode(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h handler) deleteMessage(w http.ResponseWriter, req *http.Request) {
	var del model.DeleteMessageRequest
	err := json.NewDecoder(req.Body).Decode(&del)
	if err != nil {
		http.Error(w, "Cannot Deserialize body to 'DeleteMessage' format", http.StatusBadRequest)
	}
	h.service.Delete(del.Id, del.Code)

	w.WriteHeader(http.StatusNoContent)
}
