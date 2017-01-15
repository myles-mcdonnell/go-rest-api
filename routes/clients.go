package routes

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/myles-mcdonnell/go-rest-api/logging"
	"github.com/myles-mcdonnell/go-rest-api/models"
	"github.com/myles-mcdonnell/go-rest-api/repository"
	"github.com/myles-mcdonnell/go-rest-api/restapi/operations/clients"
	"github.com/myles-mcdonnell/go-rest-api/util"
)

func ClientsGet(params clients.GetClientsParams) middleware.Responder {

	defer func() {
		logging.Log(logging.NewLogEvent(logging.DEBUG, logging.GetClientsEnd, params.HTTPRequest))
	}()

	logging.Log(logging.NewLogEvent(logging.DEBUG, logging.GetClientsStart, params.HTTPRequest))

	results, err := repository.GetAllClients()

	if err != nil {
		errMsg := &models.ErrorMessage{Message: util.SPtr(err.Error())}
		logging.Log(logging.NewLogEvent(logging.INFO, logging.GetClientsError, params.HTTPRequest).WithAdditional(errMsg))
		return clients.NewGetClientsDefault(500).WithPayload(&models.ErrorMessage{Message: util.SPtr(err.Error())})
	}

	return &clients.GetClientsOK{results}
}
