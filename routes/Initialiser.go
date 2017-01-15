package routes

import (
	"github.com/myles-mcdonnell/go-rest-api/restapi/operations"
	"github.com/myles-mcdonnell/go-rest-api/restapi/operations/clients"
	"github.com/myles-mcdonnell/go-rest-api/restapi/operations/healthcheck"
)

func BindRoutes(api *operations.GoRestAPI) {

	api.HealthcheckGetHealthcheckHandler = healthcheck.GetHealthcheckHandlerFunc(HealthcheckGet)
	api.ClientsGetClientsHandler = clients.GetClientsHandlerFunc(ClientsGet)
}
