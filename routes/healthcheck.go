package routes

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/myles-mcdonnell/go-rest-api/logging"
	"github.com/myles-mcdonnell/go-rest-api/models"
	"github.com/myles-mcdonnell/go-rest-api/repository"
	"github.com/myles-mcdonnell/go-rest-api/restapi/operations/healthcheck"
	"github.com/myles-mcdonnell/go-rest-api/util"
)

func HealthcheckGet(params healthcheck.GetHealthcheckParams) middleware.Responder {

	defer func() {
		logging.Log(logging.NewLogEvent(logging.DEBUG, logging.GetHealthcheckEnd, params.HTTPRequest))
	}()

	logging.Log(logging.NewLogEvent(logging.DEBUG, logging.GetHealthcheckStart, params.HTTPRequest))

	if err := repository.TestDbConnection(); err != nil {

		errMsg := &models.ErrorMessage{Message: util.SPtr(err.Error())}
		logging.Log(logging.NewLogEvent(logging.INFO, logging.GetHealthcheckErrorOnTestDbConnection, params.HTTPRequest).WithAdditional(errMsg))
		return healthcheck.NewGetHealthcheckDefault(500).WithPayload(errMsg)
	}

	return healthcheck.NewGetHealthcheckOK()
}
