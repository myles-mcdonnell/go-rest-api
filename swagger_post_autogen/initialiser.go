package restapi

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate/migrate"
	"github.com/myles-mcdonnell/go-rest-api/logging"
	"github.com/myles-mcdonnell/go-rest-api/repository"
	"github.com/myles-mcdonnell/go-rest-api/restapi/operations"
	"github.com/myles-mcdonnell/go-rest-api/routes"
	"github.com/myles-mcdonnell/vipermask"
	"time"
)

type gorestapiconfig struct {
	Databasename                   string
	Databasehost                   string
	Databaseport                   int
	Databasepwd                    string
	Databaseuser                   string
	Databasedisablesslmode         bool
	Databaseconnectionretryseconds int
	Maxopenconnections             int
	Maxidleconnections             int
	Serviceport                    int
	Outputconfig                   bool
	Enablecors                     bool
	LogOutputDebug                 bool
	LogOutputJson                  bool
	LogOutputPretty                bool
}

var conf *gorestapiconfig
var db *sqlx.DB

func InitApi(api *operations.GoRestAPI) {
	readConfig()

	logging.Initialise(conf.LogOutputJson, conf.LogOutputPretty, conf.LogOutputDebug)

	if conf.Outputconfig {
		logging.Log(logging.NewLogEvent(logging.INFO, logging.OutputConfig, nil).WithAdditional(conf))
	}

	api.Logger = func(msg string, args ...interface{}) {
		logging.Log(logging.NewLogEvent(logging.INFO, logging.LogEventType(msg), nil).WithAdditional(args))
	}

	initDatabase()
	runDbMigrations()
	repository.Init(db)
	routes.BindRoutes(api)
}

func readConfig() {
	confReader := config.NewFromFiles("environment.toml", "config.toml")

	conf = new(gorestapiconfig)
	conf.Databasename = confReader.GetString("database.name")
	conf.Databasehost = confReader.GetString("database.host")
	conf.Databaseport = confReader.GetInt("database.port")
	conf.Databasepwd = confReader.GetString("database.password")
	conf.Databaseuser = confReader.GetString("database.user")
	conf.Databasedisablesslmode = confReader.GetBool("database.disablesslmode")
	conf.Maxidleconnections = confReader.GetInt("database.maxidleconnections")
	conf.Maxopenconnections = confReader.GetInt("database.maxopenconnections")
	conf.Serviceport = confReader.GetInt("service.port")
	conf.Outputconfig = confReader.GetBool("service.outputconfig")
	conf.Enablecors = confReader.GetBool("service.enablecors")
	conf.LogOutputDebug = confReader.GetBool("log.outputdebug")
	conf.LogOutputPretty = confReader.GetBool("log.outputpretty")
	conf.LogOutputJson = confReader.GetBool("log.outputjson")
}

func initDatabase() {

	defer func() {
		e := recover()
		// recover from panic if one occured. Set err to nil otherwise.
		if e != nil {
			logging.Log(logging.NewLogEvent(logging.INFO, logging.InitDbError, nil).WithAdditional(fmt.Sprintf("Error occured connecting to database: %s, retry in 3 seconds", e)))
			time.Sleep(time.Second * 3)
			initDatabase()
		}
	}()

	var err error
	db, err = sqlx.Open("postgres", conf.databaseurl())
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(conf.Maxidleconnections)
	db.SetMaxOpenConns(conf.Maxopenconnections)
}

func runDbMigrations() {
	allErrors, ok := migrate.UpSync(conf.databaseurl(), "./db/migrations")
	if !ok {
		panic(allErrors[0].Error())
	}
}

func (conf *gorestapiconfig) databaseurl() string {

	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", conf.Databaseuser, conf.Databasepwd, conf.Databasehost, conf.Databaseport, conf.Databasename)

	if conf.Databasedisablesslmode {
		url = fmt.Sprintf("%s?sslmode=disable", url)
	}

	return url
}
