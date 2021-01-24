package validator

import (
	"regexp"

	"bitbucket.org/oaroz/hai/app/errors"
	"bitbucket.org/oaroz/hai/app/model"
)

var regexpEmail = regexp.MustCompile("^(.)+[@][^@]+[.][a-zA-Z0-9]+$")

func EmailFormat(email string) bool {
	if !regexpEmail.MatchString(email) {
		return false
	}
	return true
}

func CreateMessageRequest(req model.CreateMessageRequest) []string {
	errs := []string{}
	if req.Content == "" {
		errs = append(errs, errors.InvalidMessageContent)
	}
	if req.Title == "" {
		errs = append(errs, errors.InvalidMessageTitle)
	}
	if req.Email == "" {
		errs = append(errs, errors.EmptyEmail)
	} else if !regexpEmail.MatchString(req.Email) {
		errs = append(errs, errors.InvalidEmailFormat)
	}
	return errs
}

func DeleteMessageRequest(req model.DeleteMessageRequest) []string {
	errs := []string{}
	if req.Id == 0 {
		errs = append(errs, errors.RequiredFieldOf("id", "DeleteMessageRequest"))
	}
	if req.Code == "" {
		errs = append(errs, errors.EmptyFieldOf("code", "DeleteMessageRequest"))
	}
	return errs
}
