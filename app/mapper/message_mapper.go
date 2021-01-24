package mapper

import (
	"bitbucket.org/oaroz/hai/app/domain"
	"bitbucket.org/oaroz/hai/app/model"
)

func CreateReqToMessage(req model.CreateMessageRequest) domain.Message {
	var msg domain.Message
	msg.Title = req.Title
	msg.Email = req.Email
	msg.Content = req.Content
	return msg
}

func MessageToCreateResponse(msg domain.Message) model.CreateMessageResponse {
	var res model.CreateMessageResponse
	res.Id = msg.ID
	res.Code = msg.Code
	res.Email = msg.Email
	res.Title = msg.Title
	return res
}
