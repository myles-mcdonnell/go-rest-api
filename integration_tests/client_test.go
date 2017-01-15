package integration_tests

import (
	"github.com/myles-mcdonnell/go-rest-api/integration_tests/client/clients"
	"testing"
)

func Test_Client_CRUD(t *testing.T) {

	get_all_response, err := Apiclient.Clients.GetClients(clients.NewGetClientsParams())

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(get_all_response.Payload) != 0 {
		t.Fatalf("Expected zero clients")
	}
}
