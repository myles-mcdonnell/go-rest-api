package logging

import (
	"gopkg.in/myles-mcdonnell/loglight.v2"
	"gopkg.in/satori/go.uuid.v1"
	"net/http"
	"os"
	"time"
)

var logger *loglight.Logger
var hostAddress string

const requestKeyHeaderName string = "RequestKey"

var (
	GetHealthcheckStart                   LogEventType = "Get Healthcheck - Start"
	GetHealthcheckEnd                     LogEventType = "Get Healthcheck - End"
	GetClientsStart                       LogEventType = "Get All Clients - Start"
	GetClientsEnd                         LogEventType = "Get All Clients - End"
	GetClientsError                       LogEventType = "Get All Clients - Error"
	SqlStatementPrepError                 LogEventType = "Sql Statment Prep Error"
	GetHealthcheckErrorOnTestDbConnection LogEventType = "Get Healthcheck Error On Test Db §Connection"
	OutputConfig                          LogEventType = "Output Config"
	InitDbError                           LogEventType = "Init Db Error"
)

func Initialise(outputJson bool, outputPretty bool, outputDebug bool) {

	logger = loglight.NewLoggerWithOptions(outputDebug, 0, outputJson, outputPretty)
	hostAddress, _ = os.Hostname()
}

type LogEvent struct {
	TimeUtc     time.Time
	ServiceKey  string
	Title       LogEventType
	Additional  interface{}
	RequestKey  string
	HostAddress string
	Level       LogLevel
}

type LogEventType string
type LogLevel string

var (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	ERROR LogLevel = "ERROR"
)

func Log(event *LogEvent) {
	switch event.Level {

	case DEBUG:
		logger.LogDebugStruct(event)
	case INFO:
		logger.LogInfoStruct(event)
	case ERROR:
		logger.LogInfoStruct(event)
	}
}

func NewLogEvent(logLevel LogLevel, logType LogEventType, request *http.Request) *LogEvent {

	var requestKey string = ""

	if request != nil {
		requestKey = request.Header.Get(requestKeyHeaderName)
	}

	return &LogEvent{
		ServiceKey:  "GORESTAPI_SVC",
		TimeUtc:     time.Now().UTC(),
		Title:       logType,
		RequestKey:  requestKey,
		HostAddress: hostAddress,
		Level:       logLevel,
	}
}

func (logEvent *LogEvent) WithAdditional(additional interface{}) *LogEvent {
	logEvent.Additional = additional
	return logEvent
}

type RequestKey struct {
	handler http.Handler
}

func (key *RequestKey) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(requestKeyHeaderName) == "" {
		r.Header.Set(requestKeyHeaderName, uuid.NewV4().String())
	}
}

// Handler apply the CORS specification on the request, and add relevant CORS headers
// as necessary.
func (c *RequestKey) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c.ServeHTTP(w, r)
		h.ServeHTTP(w, r)
	})
}
