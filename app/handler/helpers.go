package handler

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/oaroz/hai/app/errors"
)

func HandleErrors(err errors.HaiError, w http.ResponseWriter) {
	switch err.ErrorCode {
	case errors.MessageNotFound:
		ReturnJson(w, errors.NewErrResponse([]string{errors.MessageNotFound}, []string{errors.MessageNotFoundDetails}), http.StatusNotFound)
	case errors.DeleteNonexistentMessage:
		ReturnJson(w, errors.NewErrResponse([]string{errors.DeleteNonexistentMessage}, []string{errors.DeleteNonexistentMessageDetails}), http.StatusBadRequest)
	case errors.DatabaseActionError:
		ReturnJson(w, errors.NewErrResponse([]string{errors.DatabaseActionError}, []string{}), http.StatusInternalServerError)
	case errors.InvalidBodyRequest:
		ReturnJson(w, errors.NewErrResponse([]string{errors.InvalidBodyRequest}, []string{}), http.StatusBadRequest)
	case errors.InvalidQueryParam:
		ReturnJson(w, errors.NewErrResponse([]string{errors.InvalidQueryParam}, []string{}), http.StatusBadRequest)
	default:
		ReturnJson(w, errors.NewErrResponse([]string{errors.InternalServerError}, []string{errors.InternalServerErrorDetails}), http.StatusInternalServerError)
	}
}

func ReturnJson(w http.ResponseWriter, out interface{}, httpCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	if out != nil {
		json.NewEncoder(w).Encode(out)
	}
}
