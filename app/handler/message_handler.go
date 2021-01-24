package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/oaroz/hai/app/errors"
	"bitbucket.org/oaroz/hai/app/mapper"
	"bitbucket.org/oaroz/hai/app/model"
	"bitbucket.org/oaroz/hai/app/validator"

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
	if !validator.Email(email) {
		http.Error(w, "Passed 'email' query param in the wrong format.", http.StatusBadRequest)
		return
	}
	messages, err := h.service.Get(email)
	if err != nil {
		fmt.Println(err.(errors.HaiError))
		fmt.Println(err.(errors.HaiError).Message)
	}
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
	if !validator.CreateMessageRequest(msgReq) {
		http.Error(w, "Passed 'request body' in the wrong format.", http.StatusBadRequest)
		return
	}
	msg, err := h.service.Create(mapper.CreateReqToMessage(msgReq))
	if err != nil {
		fmt.Println(err.(errors.HaiError))
		fmt.Println(err.(errors.HaiError).Message)
	}
	var res model.CreateMessageResponse = mapper.MessageToCreateResponse(msg)
	json.NewEncoder(w).Encode(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h handler) deleteMessage(w http.ResponseWriter, req *http.Request) {
	var del model.DeleteMessageRequest
	err := json.NewDecoder(req.Body).Decode(&del)
	if err != nil {
		http.Error(w, "Cannot Deserialize body to 'DeleteMessage' format", http.StatusBadRequest)
		return
	}
	err = h.service.Delete(del.Id, del.Code)
	if err != nil {
		fmt.Println(err.(errors.HaiError))
		fmt.Println(err.(errors.HaiError).Message)
	}
	w.WriteHeader(http.StatusNoContent)
}
