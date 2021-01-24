package validator

import (
	"regexp"

	"bitbucket.org/oaroz/hai/app/model"
)

var regexpEmail = regexp.MustCompile("^(.)+[@][^@]+[.][a-zA-Z0-9]+$")

func Email(email string) bool {
	if email == "" {
		return true
	}

	if !regexpEmail.MatchString(email) {
		return false
	}
	return true
}

func CreateMessageRequest(req model.CreateMessageRequest) bool {
	if req.Content == "" {
		return false
	}
	if len(req.Title) < 2 {
		return false
	}
	if req.Email == "" {
		return false
	}
	if !regexpEmail.MatchString(req.Email) {
		return false
	}
	return true
}
