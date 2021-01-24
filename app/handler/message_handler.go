package handler

import (
	"encoding/json"
	"log"
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

	if email != "" {
		if !validator.EmailFormat(email) {
			ReturnJson(
				w,
				errors.NewErrResponse([]string{errors.InvalidQueryParam}, []string{errors.InvalidEmailFormat}),
				http.StatusBadRequest)
			return
		}
	}
	messages, err := h.service.Get(email)

	if err != nil {
		HandleErrors(err.(errors.HaiError), w)
		return
	}
	ReturnJson(w, messages, http.StatusOK)
}

func (h handler) postMessage(w http.ResponseWriter, req *http.Request) {
	var msgReq model.CreateMessageRequest
	err := json.NewDecoder(req.Body).Decode(&msgReq)

	if err != nil {
		log.Printf("Cannot Deserialize. Invalid json. %v", err)
		ReturnJson(w, errors.NewErrResponse([]string{errors.InvalidBodyRequest}, []string{errors.InvalidJson}), http.StatusBadRequest)
		return
	}
	errs := validator.CreateMessageRequest(msgReq)

	if len(errs) != 0 {
		ReturnJson(w, errors.NewErrResponse([]string{errors.InvalidBodyRequest}, errs), http.StatusBadRequest)
		return
	}
	msg, err := h.service.Create(mapper.CreateReqToMessage(msgReq))
	var res model.CreateMessageResponse = mapper.MessageToCreateResponse(msg)

	if err != nil {
		HandleErrors(err.(errors.HaiError), w)
		return
	}
	ReturnJson(w, res, http.StatusCreated)
}

func (h handler) deleteMessage(w http.ResponseWriter, req *http.Request) {
	var del model.DeleteMessageRequest
	err := json.NewDecoder(req.Body).Decode(&del)

	if err != nil {
		log.Printf("Cannot Deserialize. Invalid json. %v", err)
		ReturnJson(w, errors.NewErrResponse([]string{errors.InvalidBodyRequest}, []string{errors.InvalidJson}), http.StatusBadRequest)
		return
	}
	errs := validator.DeleteMessageRequest(del)

	if len(errs) != 0 {
		ReturnJson(w, errors.NewErrResponse([]string{errors.InvalidBodyRequest}, errs), http.StatusBadRequest)
		return
	}
	err = h.service.Delete(del.Id, del.Code)

	if err != nil {
		HandleErrors(err.(errors.HaiError), w)
		return
	}
	ReturnJson(w, nil, http.StatusNoContent)
}
