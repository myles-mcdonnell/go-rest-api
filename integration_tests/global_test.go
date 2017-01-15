package integration_tests

import (
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/myles-mcdonnell/go-rest-api/integration_tests/client"
	"github.com/myles-mcdonnell/go-rest-api/integration_tests/client/healthcheck"
	"gopkg.in/myles-mcdonnell/vipermask.v1"
	"log"
	"os"
	"testing"
	"time"
)

type testconfig struct {
	servicehostandport string
}

var Apiclient *client.Gorestapi
var conf *testconfig
var initserviceretrycount int

func TestMain(m *testing.M) {

	readConfig()

	httptrans := httptransport.New(conf.servicehostandport, "/", []string{"http"})
	Apiclient = client.New(httptrans, nil)

	if establishconnection() {
		os.Exit(m.Run())
	}

	os.Exit(1)
}

func establishconnection() bool {
	defer func() {

		e := recover()
		// recover from panic if one occurred. Set err to nil otherwise.
		if e != nil {
			log.Printf("init attempt %d failed. Retry in %d seconds", initserviceretrycount, 1)
			time.Sleep(time.Second * 1)
			initserviceretrycount++
			establishconnection()
		}
	}()

	log.Printf("Connecting to go rest api Service : %s", conf.servicehostandport)

	_, err := Apiclient.Healthcheck.GetHealthcheck(healthcheck.NewGetHealthcheckParams())

	if err != nil {
		log.Printf("Error: %s", err.Error())
		if initserviceretrycount > 9 {
			return false
		}
		panic(err.Error())
	}

	log.Printf("Connected to service OK : %s", conf.servicehostandport)

	return true
}

func SPtr(s string) *string { return &s }

func (conf *testconfig) servicebaseurl() string {

	return fmt.Sprintf("http://%s", conf.servicehostandport)
}

func readConfig() {

	confReader := config.NewFromFiles("config.environment.toml", "config.toml")

	conf = new(testconfig)
	conf.servicehostandport = fmt.Sprintf("%s:%d", confReader.GetString("service.apihost"), confReader.GetInt("service.apiport"))
}
